package main

import (
	"bytes"
	"io"
	"math"
	"reflect"
	"testing"
)

type TestErrorWriter struct{}

func (testWriter *TestErrorWriter) Write(p []byte) (n int, err error) {
	return 0, io.EOF
}

func TestCpu_MoveRight(t *testing.T) {
	memory := make([]uint8, 3)
	cpu := Cpu{memory: memory}

	cpu.Execute(Instruction{name: MoveRight})

	if cpu.addressCounter != 1 {
		t.Errorf("unexpected address counter, got %d, want %d", cpu.addressCounter, 1)
	}

	cpu.Execute(Instruction{name: MoveRight})

	if cpu.addressCounter != 2 {
		t.Errorf("unexpected address counter, got %d, want %d", cpu.addressCounter, 2)
	}

	cpu.Execute(Instruction{name: MoveRight})

	if cpu.addressCounter != 0 {
		t.Errorf("unexpected address counter, got %d, want %d", cpu.addressCounter, 0)
	}
}

func TestCpu_MoveLeft(t *testing.T) {
	memory := make([]uint8, 3)
	cpu := Cpu{memory: memory, addressCounter: 2}

	cpu.Execute(Instruction{name: MoveLeft})

	if cpu.addressCounter != 1 {
		t.Errorf("unexpected address counter, got %d, want %d", cpu.addressCounter, 1)
	}

	cpu.Execute(Instruction{name: MoveLeft})

	if cpu.addressCounter != 0 {
		t.Errorf("unexpected address counter, got %d, want %d", cpu.addressCounter, 0)
	}

	cpu.Execute(Instruction{name: MoveLeft})

	if cpu.addressCounter != 2 {
		t.Errorf("unexpected address counter, got %d, want %d", cpu.addressCounter, 2)
	}
}

func TestCpu_IncrementInstruction(t *testing.T) {
	memory := make([]uint8, 2)
	cpu := Cpu{memory: memory}

	cpu.Execute(Instruction{name: Increment})

	if memory[0] != 1 {
		t.Errorf("unexpected memory state, got %d, want %d", memory[0], 1)
	}

	cpu.Execute(Instruction{name: Increment})

	if memory[0] != 2 {
		t.Errorf("unexpected memory state, got %d, want %d", memory[0], 2)
	}

	cpu.Execute(Instruction{name: Increment})

	if memory[0] != 3 {
		t.Errorf("unexpected memory state, got %d, want %d", memory[0], 3)
	}
}

func TestCpu_DecrementInstruction(t *testing.T) {
	memory := []uint8{3}
	cpu := Cpu{memory: memory}

	cpu.Execute(Instruction{name: Decrement})

	if memory[0] != 2 {
		t.Errorf("unexpected memory state, got %d, want %d", memory[0], 2)
	}

	cpu.Execute(Instruction{name: Decrement})

	if memory[0] != 1 {
		t.Errorf("unexpected memory state, got %d, want %d", memory[0], 1)
	}

	cpu.Execute(Instruction{name: Decrement})

	if memory[0] != 0 {
		t.Errorf("unexpected memory state, got %d, want %d", memory[0], 0)
	}
}

func TestCpu_WriteInstruction(t *testing.T) {
	memory := []uint8{97, 122} // a, z
	var output bytes.Buffer
	cpu := Cpu{memory: memory, output: &output}

	cpu.Execute(Instruction{name: Write})

	byte, err := output.ReadByte()
	if err != nil {
		t.Fatal(err)
	}

	if byte != 97 {
		t.Errorf("unexpected write value, got %c, want %c", byte, 97)
	}

	cpu.Execute(Instruction{name: MoveRight})
	cpu.Execute(Instruction{name: Write})

	byte, err = output.ReadByte()
	if err != nil {
		t.Fatal(err)
	}

	if byte != 122 {
		t.Errorf("unexpected write value, got %c, want %c", byte, 97)
	}

	defer func() {
		if err := recover(); err == io.EOF {
			t.Errorf("read instruction code did not panic")
		}
	}()

	var panicOutput TestErrorWriter
	cpu.output = &panicOutput

	cpu.Execute(Instruction{name: Write})
}

func TestCpu_ReadInstruction(t *testing.T) {
	memory := make([]uint8, 2)
	var input bytes.Buffer
	cpu := Cpu{memory: memory, input: &input}

	err := input.WriteByte(97) // a
	if err != nil {
		t.Fatal(err)
	}

	cpu.Execute(Instruction{name: Read})

	if memory[0] != 97 {
		t.Errorf("unexpected memory state, got %c, want %c", memory[0], 97)
	}

	cpu.Execute(Instruction{name: MoveRight})

	err = input.WriteByte(122) // z
	if err != nil {
		t.Fatal(err)
	}

	cpu.Execute(Instruction{name: Read})

	if memory[1] != 122 {
		t.Errorf("unexpected memory state, got %c, want %c", memory[1], 122)
	}

	defer func() {
		if err := recover(); err == io.EOF {
			t.Errorf("read instruction code did not panic")
		}
	}()

	cpu.Execute(Instruction{name: Read}) // no input to force EOF error
}

func TestCpu_JumpIfZeroInstruction(t *testing.T) {
	memory := make([]uint8, 1)
	cpu := Cpu{memory: memory}

	// emulated instructions: [ ++ ]

	cpu.Execute(Instruction{name: JumpUnlessZero, linkedJump: 4})

	if cpu.programCounter != 1 {
		t.Errorf("unexpected program counter, got %d, want %d", cpu.programCounter, 1)
	}

	cpu.Execute(Instruction{name: Increment})
	cpu.Execute(Instruction{name: Increment})

	if memory[0] != 2 {
		t.Errorf("unexpected memory state, got %d, want %d", memory[0], 2)
	}

	cpu.Execute(Instruction{name: JumpIfZero, linkedJump: 0})

	if cpu.programCounter != 4 {
		t.Errorf("unexpected program counter, got %d, want %d", cpu.programCounter, 4)
	}
}

func TestCpu_JumpUnlessZeroInstruction(t *testing.T) {
	memory := make([]uint8, 1)
	cpu := Cpu{memory: memory}

	// emulated instructions: [ + ] [ - ]

	cpu.Execute(Instruction{name: JumpUnlessZero, linkedJump: 2})

	if cpu.programCounter != 1 {
		t.Errorf("unexpected program counter, got %d, want %d", cpu.programCounter, 1)
	}

	cpu.Execute(Instruction{name: Increment})

	if memory[0] != 1 {
		t.Errorf("unexpected memory state, got %d, want %d", memory[0], 1)
	}

	cpu.Execute(Instruction{name: JumpIfZero, linkedJump: 0})

	if cpu.programCounter != 3 {
		t.Errorf("unexpected program counter, got %d, want %d", cpu.programCounter, 3)
	}

	cpu.Execute(Instruction{name: JumpUnlessZero, linkedJump: 5})

	if cpu.programCounter != 5 {
		t.Errorf("unexpected program counter, got %d, want %d", cpu.programCounter, 5)
	}

	cpu.Execute(Instruction{name: Decrement})

	if memory[0] != 0 {
		t.Errorf("unexpected memory state, got %d, want %d", memory[0], 0)
	}

	cpu.Execute(Instruction{name: JumpIfZero, linkedJump: 3})

	if cpu.programCounter != 3 {
		t.Errorf("unexpected program counter, got %d, want %d", cpu.programCounter, 5)
	}
}

func TestCpu_ExecuteSimple(t *testing.T) {
	memory := make([]uint8, 2)
	instructions := []Instruction{
		{name: Increment},
		{name: Increment},
		{name: MoveRight},
		{name: Increment},
		{name: MoveLeft},
		{name: Decrement},
	}

	cpu := Cpu{
		memory: memory,
		code:   instructions,
	}

	if !reflect.DeepEqual(memory, []uint8{0, 0}) {
		t.Errorf("unexpected memory state, got %v, want %v", memory, []uint8{0, 0})
	}

	cpu.Step()

	if !reflect.DeepEqual(memory, []uint8{1, 0}) {
		t.Errorf("unexpected memory state, got %v, want %v", memory, []uint8{1, 0})
	}

	cpu.Step()

	if !reflect.DeepEqual(memory, []uint8{2, 0}) {
		t.Errorf("unexpected memory state, got %v, want %v", memory, []uint8{2, 0})
	}

	cpu.Step()

	if !reflect.DeepEqual(memory, []uint8{2, 0}) {
		t.Errorf("unexpected memory state, got %v, want %v", memory, []uint8{2, 0})
	}

	cpu.Step()

	if !reflect.DeepEqual(memory, []uint8{2, 1}) {
		t.Errorf("unexpected memory state, got %v, want %v", memory, []uint8{2, 1})
	}

	cpu.Step()

	if !reflect.DeepEqual(memory, []uint8{2, 1}) {
		t.Errorf("unexpected memory state, got %v, want %v", memory, []uint8{2, 1})
	}

	cpu.Step()

	if !reflect.DeepEqual(memory, []uint8{1, 1}) {
		t.Errorf("unexpected memory state, got %v, want %v", memory, []uint8{1, 1})
	}
}

func TestCpu_HasInstructionsLeft(t *testing.T) {
	memory := make([]uint8, 1)
	instructions := []Instruction{
		{name: Increment},
		{name: Increment},
	}

	cpu := Cpu{
		memory: memory,
		code:   instructions,
	}

	if !cpu.HasInstructionsLeft() {
		t.Errorf("unexpected state, cpu has no instructions left")
	}

	cpu.Step()

	if !cpu.HasInstructionsLeft() {
		t.Errorf("unexpected state, cpu has no instructions left")
	}

	cpu.Step()

	if cpu.HasInstructionsLeft() {
		t.Errorf("unexpected state, cpu has instructions left")
	}
}

func TestCpu_MemoryOverflow(t *testing.T) {
	memory := make([]uint8, 1)

	cpu := Cpu{memory: memory}

	for i := 0; i < math.MaxUint8; i++ {
		cpu.Execute(Instruction{name: Increment})
	}

	if memory[0] != math.MaxUint8 {
		t.Errorf("unexpected memory state, got %d, want %d", memory[0], math.MaxUint8)
	}

	cpu.Execute(Instruction{name: Increment})

	if memory[0] != 0 {
		t.Errorf("unexpected memory state, got %d, want %d", memory[0], 0)
	}
}

func TestCpu_MemoryUnderflow(t *testing.T) {
	memory := make([]uint8, 1)

	cpu := Cpu{memory: memory}

	cpu.Execute(Instruction{name: Decrement})

	if memory[0] != math.MaxUint8 {
		t.Errorf("unexpected memory state, got %d, want %d", memory[0], math.MaxUint8)
	}
}
