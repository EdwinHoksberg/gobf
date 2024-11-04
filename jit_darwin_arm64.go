//go:build darwin && arm64

package main

import (
	"encoding/binary"
	"errors"
	"syscall"
	"unsafe"
)

func (jit *Jit) compile(instructions []Instruction) error {
	// x0 contains a pointer to program memory
	// x1 contains a pointer to executable memory

	// x9 = address counter
	// x10 = program counter (can be removed?)
	// x11 = scratch
	// x15 = pointer to program memory

	jit.code = append(jit.code,
		// reset registers x9, x10, x11 to 0
		0x09, 0x00, 0x80, 0xd2, // mov x9, #0
		0x0a, 0x00, 0x80, 0xd2, // mov x10, #0
		0x0b, 0x00, 0x80, 0xd2, // mov x11, #0

		// move first argument(pointer to program memory) to x15
		0xef, 0x03, 0x00, 0xaa, // mov x15, x0

		// move second argument(pointer to executable memory) to x14
		0xee, 0x03, 0x01, 0xaa, // mov x14, x1
	)

	for _, instruction := range instructions {
		block := CodeBlock{
			instruction: instruction,
			offset:      len(jit.code),
		}

		switch instruction.name {
		case MoveRight:
			// @todo check for overflow(> size of program memory) and wrap
			jit.code = append(jit.code,
				// increase the address counter by one
				0x29, 0x05, 0x00, 0x91, // add x9, x9, #1
			)
		case MoveLeft:
			// @todo check for underflow(< 0) and wrap
			jit.code = append(jit.code,
				// decrease the address counter by one
				0x29, 0x05, 0x00, 0xd1, // sub x9, x9, #1
			)
		case Increment:
			jit.code = append(jit.code,
				// load the current value of the program memory offset by the address counter
				0xeb, 0x69, 0x69, 0x38, // ldrb w11, [x15, x9]

				// add one to the value which we've loaded
				0x6b, 0x05, 0x00, 0x91, // add x11, x11, #1

				// store the value back to the program memory including offset
				0xeb, 0x69, 0x29, 0x38, // strb w11, [x15, x9]
			)
		case Decrement:
			jit.code = append(jit.code,
				// load the current value of the program memory offset by the address counter
				0xeb, 0x69, 0x69, 0x38, // ldrb w11, [x15, x9]

				// subtract one to the value which we've loaded
				0x6b, 0x05, 0x00, 0xd1, // sub x11, x11, #1

				// store the value back to the program memory including offset
				0xeb, 0x69, 0x29, 0x38, // strb w11, [x15, x9]
			)
		case Write:
			jit.code = append(jit.code,
				// arg 1, stdout file descriptor
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
		case Read:
			jit.code = append(jit.code,
				// arg 1, stdout file descriptor
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
		case JumpIfZero:
			jit.code = append(jit.code,
				// load the current value of the program memory offset by the address counter
				0xeb, 0x69, 0x69, 0x38, // ldrb w11, [x15, x9]

				// jump to right before the linked jump instruction
				0x2b, 0x00, 0x00, 0x34, // cbz w11, #0
			)
		case JumpUnlessZero:
			jit.code = append(jit.code,
				// load the current value of the program memory offset by the address counter
				0xeb, 0x69, 0x69, 0x38, // ldrb w11, [x15, x9]

				// jump to right after the linked jump instruction
				0x00, 0x00, 0x00, 0x35, // cbnz w11, #0
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

	jit.postProcessJumps()

	return nil
}

func (jit *Jit) run(memorySize uint) error {
	// Allocate program memory
	programMemory, err := syscall.Mmap(-1, 0, int(memorySize), syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_ANON|syscall.MAP_PRIVATE)
	if err != nil {
		return errors.New("failed to map program memory: " + err.Error())
	}

	// Allocate executable memory
	executableMemory, err := syscall.Mmap(-1, 0, len(jit.code), syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_ANON|syscall.MAP_PRIVATE)
	if err != nil {
		return errors.New("failed to map executable memory: " + err.Error())
	}

	// Copy JIT instructions to executable memory
	copy(executableMemory, jit.code)

	// Change permissions to executable
	if err := syscall.Mprotect(executableMemory, syscall.PROT_EXEC); err != nil {
		return errors.New("failed to make executable memory executable: " + err.Error())
	}

	// Cast allocated memory regions to function pointers
	executableMemoryPointer := &executableMemory
	programMemoryPointer := unsafe.Pointer(&programMemory[0])

	// Define JIT call function and execute it
	f := *(*func(programMemory unsafe.Pointer, executableMemory unsafe.Pointer))(unsafe.Pointer(&executableMemoryPointer))
	f(programMemoryPointer, unsafe.Pointer(&executableMemory[0]))

	// Flush any changes made by the JIT code to program memory
	if _, _, errno := syscall.Syscall(syscall.SYS_MSYNC, uintptr(programMemoryPointer), uintptr(memorySize), syscall.MS_SYNC); errno != 0 {
		return errors.New("failed to sync program memory after running: " + errno.Error())
	}

	return nil
}

func (jit *Jit) postProcessJumps() {
	for i, block := range jit.codeBlocks {
		if block.instruction.name != JumpIfZero && block.instruction.name != JumpUnlessZero {
			continue
		}

		jit.codeBlocks[i].link = &jit.codeBlocks[block.instruction.link]
	}

	for _, block := range jit.codeBlocks {
		if block.instruction.name != JumpIfZero && block.instruction.name != JumpUnlessZero {
			// Only process jump instructions
			continue
		}

		if block.link == nil {
			// Only process linked blocks
			continue
		}

		offset := (block.link.offset - block.offset + 4) / 4

		// Sign-extend the offset to 19 bits
		if offset&(1<<18) != 0 { // If the 19th bit (sign bit) is set
			offset |= ^((1 << 19) - 1) // Extend the sign bit to 32 bits
		}

		// Ensure the offset fits in a signed 19-bit integer (-2^18 to 2^18 - 1)
		if offset < -(1<<18) || offset >= (1<<18) {
			panic("offset is out of range for a jump")
		}

		// Base opcode for CBZ
		opcode := uint32(0x34000000)
		if block.instruction.name == JumpUnlessZero {
			// Base opcode for CBNZ
			opcode = uint32(0x35000000)
		}

		// Encode the offset into bits [23:5]
		opcode |= (uint32(offset) & 0x7FFFF) << 5

		// Encode the register into bits [4:0]
		opcode |= uint32(11 & 0x1F) // 11 = w11

		// +4 because we need to insert it in after the ldrb instruction
		binary.LittleEndian.PutUint32(jit.code[block.offset+4:], opcode)
	}
}
