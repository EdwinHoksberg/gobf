package main

import (
	"fmt"
	"io"
)

type Cpu struct {
	memory []uint8
	code   []Instruction
	input  io.Reader
	output io.Writer

	programCounter int
	addressCounter int
}

func (cpu *Cpu) Step() {
	cpu.Execute(cpu.code[cpu.programCounter])
}

func (cpu *Cpu) Execute(instruction Instruction) {
	switch instruction.name {
	case MoveRight:
		if cpu.addressCounter == len(cpu.memory)-1 {
			cpu.addressCounter = 0
		} else {
			cpu.addressCounter++
		}
	case MoveLeft:
		if cpu.addressCounter == 0 {
			cpu.addressCounter = len(cpu.memory) - 1
		} else {
			cpu.addressCounter--
		}
	case Increment:
		cpu.memory[cpu.addressCounter]++
	case Decrement:
		cpu.memory[cpu.addressCounter]--
	case Write:
		_, err := cpu.output.Write([]byte{cpu.memory[cpu.addressCounter]})

		if err != nil {
			panic(fmt.Sprintf("unrecoverable io error: %s", err))
		}
	case Read:
		data := make([]byte, 1)
		bytesRead, err := cpu.input.Read(data)

		if err != nil {
			panic(fmt.Sprintf("unrecoverable io error: %s", err))
		}

		if bytesRead > 0 {
			cpu.memory[cpu.addressCounter] = data[0]
		}
	case JumpIfZero:
		if cpu.memory[cpu.addressCounter] == 0 {
			cpu.programCounter = instruction.jumpPoint - 1
		}
	case JumpUnlessZero:
		if cpu.memory[cpu.addressCounter] != 0 {
			cpu.programCounter = instruction.jumpPoint - 1
		}
	}

	cpu.programCounter++
}

func (cpu *Cpu) HasInstructionsLeft() bool {
	return cpu.programCounter < len(cpu.code)
}
