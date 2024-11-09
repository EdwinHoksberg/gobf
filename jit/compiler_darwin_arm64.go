//go:build darwin && arm64

package jit

import (
	"encoding/binary"
	"errors"
	"gobf/instructions"
)

const (
	OpcodeAdd  = uint32(0x91000000)
	OpcodeSub  = uint32(0xd1000000)
	OpcodeCbz  = uint32(0x34000000)
	OpcodeCbnz = uint32(0x35000000)
)

func (jit *Jit) Compile(parsedInstructions []instructions.Instruction) error {
	// x0 contains a pointer to program memory
	// x1 contains a pointer to executable memory

	// x9 = address counter
	// x10 = program counter
	// x11 = scratch
	// x15 = pointer to program memory

	jit.code = append(jit.code,
		// reset registers x9, x10, x11 to 0
		0x09, 0x00, 0x80, 0xd2, // mov x9, #0
		0x0a, 0x00, 0x80, 0xd2, // mov x10, #0
		0x0b, 0x00, 0x80, 0xd2, // mov x11, #0

		// move first argument(pointer to program memory) to x15
		0xef, 0x03, 0x00, 0xaa, // mov x15, x0
	)

	for _, instruction := range parsedInstructions {
		block := CodeBlock{
			instruction: instruction,
			offset:      len(jit.code),
		}

		switch instruction.Name {
		case instructions.MoveRight:
			// increase the address counter by one
			if err := jit.encodeAndAppendMathInstruction(OpcodeAdd, 9, instruction.Value); err != nil {
				return err
			}
		case instructions.MoveLeft:
			// decrease the address counter by instruction value
			if err := jit.encodeAndAppendMathInstruction(OpcodeSub, 9, instruction.Value); err != nil {
				return err
			}
		case instructions.Increment:
			// load the current value of the program memory offset by the address counter
			jit.code = append(jit.code, 0xeb, 0x69, 0x69, 0x38) // ldrb w11, [x15, x9]

			// add instruction value to the value which we've loaded
			if err := jit.encodeAndAppendMathInstruction(OpcodeAdd, 11, instruction.Value); err != nil {
				return err
			}

			// store the value back to the program memory including offset
			jit.code = append(jit.code, 0xeb, 0x69, 0x29, 0x38) // strb w11, [x15, x9]
		case instructions.Decrement:
			// load the current value of the program memory offset by the address counter
			jit.code = append(jit.code, 0xeb, 0x69, 0x69, 0x38) // ldrb w11, [x15, x9]

			// subtract the instruction value from the value which we've loaded
			if err := jit.encodeAndAppendMathInstruction(OpcodeSub, 11, instruction.Value); err != nil {
				return err
			}

			// store the value back to the program memory including offset
			jit.code = append(jit.code, 0xeb, 0x69, 0x29, 0x38) // strb w11, [x15, x9]
		case instructions.Write:
			jit.code = append(jit.code,
				// arg 1, file descriptor, 1 = stdout
				0x20, 0x00, 0x80, 0xd2, // mov x0, #1

				// arg 2, pointer to code string
				0xe1, 0x03, 0x0f, 0xaa, // mov x1, x15

				// add address counter as offset to pointer to code string (arg 2)
				0x21, 0x00, 0x09, 0x8b, // add x1, x1, x9

				// arg 3, length to print, always 1
				0x22, 0x00, 0x80, 0xd2, // mov x2, #1

				// arg 4, syscall number, 4 = write
				0x90, 0x00, 0x80, 0xd2, // mov x16, #4

				// execute syscall
				0x01, 0x10, 0x00, 0xd4, // svc #0x80
			)
		case instructions.Read:
			jit.code = append(jit.code,
				// arg 1, file descriptor, 0 = stdin
				0x00, 0x00, 0x80, 0xd2, // mov x0, #0

				// arg 2, pointer to input buffer
				0xe1, 0x03, 0x0f, 0xaa, // mov x1, x15

				// add address counter as offset to pointer to input buffer (arg 2)
				0x21, 0x00, 0x09, 0x8b, // add x1, x1, x9

				// arg 3, length to read, always 1
				0x22, 0x00, 0x80, 0xd2, // mov x2, #1

				// arg 4, syscall number, 3 = read
				0x70, 0x00, 0x80, 0xd2, // mov x16, #3

				// execute syscall
				0x01, 0x10, 0x00, 0xd4, // svc #0x80
			)
		case instructions.JumpIfZero:
			jit.code = append(jit.code,
				// load the current value of the program memory offset by the address counter
				0xeb, 0x69, 0x69, 0x38, // ldrb w11, [x15, x9]

				// jump to right before the linked jump instruction
				0x0, 0x0, 0x0, 0x0, // placeholder
			)
		case instructions.JumpUnlessZero:
			jit.code = append(jit.code,
				// load the current value of the program memory offset by the address counter
				0xeb, 0x69, 0x69, 0x38, // ldrb w11, [x15, x9]

				// jump to right after the linked jump instruction
				0x0, 0x0, 0x0, 0x0, // placeholder
			)
		}

		jit.codeBlocks = append(jit.codeBlocks, block)
	}

	jit.code = append(jit.code,
		// move program memory pointer back to x0, so its used as the return value
		0xe0, 0x03, 0x0f, 0xaa, // mov x0, x15

		// return back to our Go program
		0xc0, 0x03, 0x5f, 0xd6, // ret
	)

	if err := jit.postProcessJumps(); err != nil {
		return err
	}

	return nil
}

func (jit *Jit) postProcessJumps() error {
	for i, block := range jit.codeBlocks {
		if !block.instruction.IsJump() {
			continue
		}

		jit.codeBlocks[i].link = &jit.codeBlocks[block.instruction.Value]
	}

	for _, block := range jit.codeBlocks {
		// Only process jump instructions
		if !block.instruction.IsJump() {
			continue
		}

		// Only process linked blocks
		if block.link == nil {
			return errors.New("failed to link code block")
		}

		opcode := OpcodeCbz
		if block.instruction.Name == instructions.JumpUnlessZero {
			opcode = OpcodeCbnz
		}

		// Offset is the linked block offset minus our own offset
		offset := block.link.offset - block.offset + 4

		opcode, err := encodeBranchInstruction(opcode, 11, offset)
		if err != nil {
			return err
		}

		// +4 because we need to insert it in after the ldrb instruction
		binary.LittleEndian.PutUint32(jit.code[block.offset+4:], opcode)
	}

	return nil
}

func (jit *Jit) encodeAndAppendMathInstruction(opcode uint32, register int, immediate int) error {
	if immediate < 0 || immediate >= 1<<12 {
		return errors.New("immediate out of range")
	}

	// Encode imm12 (bits 21:10)
	opcode |= uint32(immediate) << 10

	// Encode Rn (bits 9:5) and Rd (bits 4:0)
	opcode |= uint32(register&0x1F) << 5
	opcode |= uint32(register & 0x1F)

	jit.code = binary.LittleEndian.AppendUint32(jit.code, opcode)

	return nil
}

func encodeBranchInstruction(opcode uint32, register int, offset int) (uint32, error) {
	// Divide by 4 since instructions are always 4 bytes in length
	offset /= 4

	// Sign-extend the offset to 19 bits
	if offset&(1<<18) != 0 { // If the 19th bit (sign bit) is set
		offset |= ^((1 << 19) - 1) // Extend the sign bit to 32 bits
	}

	// Ensure the offset fits in a signed 19-bit integer (-2^18 to 2^18 - 1)
	if offset < -(1<<18) || offset >= (1<<18) {
		return 0, errors.New("offset is out of range for a jump")
	}

	// Encode the offset into bits [23:5]
	opcode |= (uint32(offset) & 0x7FFFF) << 5

	// Encode the register into bits [4:0]
	opcode |= uint32(register & 0x1F)

	return opcode, nil
}
