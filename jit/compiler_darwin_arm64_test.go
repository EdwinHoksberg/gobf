package jit

import (
	"bytes"
	"gobf/instructions"
	"testing"
)

func TestJit_Compile(t *testing.T) {
	var testInstructions = []instructions.Instruction{
		{Name: instructions.MoveRight},
		{Name: instructions.MoveLeft},
		{Name: instructions.Increment},
		{Name: instructions.JumpIfZero},
		{Name: instructions.Increment},
		{Name: instructions.Decrement},
		{Name: instructions.Increment},
		{Name: instructions.JumpUnlessZero},
		{Name: instructions.Read},
		{Name: instructions.Write},
	}

	jit := NewJit(1000)
	err := jit.Compile(testInstructions)

	if err != nil {
		t.Errorf("expected no error, got error: %s", err)
	}

	if !bytes.Equal(jit.code, []byte{
		0x9, 0x0, 0x80, 0xD2, 0xA, 0x0, 0x80, 0xD2, 0xB, 0x0, 0x80, 0xD2, 0xEF, 0x3, 0x0, 0xAA, 0x29, 0x5, 0x0, 0x91,
		0x29, 0x5, 0x0, 0xD1, 0xEB, 0x69, 0x69, 0x38, 0x6B, 0x5, 0x0, 0x91, 0xEB, 0x69, 0x29, 0x38, 0xEB, 0x69, 0x69,
		0x38, 0x8B, 0xFF, 0xFF, 0x34, 0xEB, 0x69, 0x69, 0x38, 0x6B, 0x5, 0x0, 0x91, 0xEB, 0x69, 0x29, 0x38, 0xEB, 0x69,
		0x69, 0x38, 0x6B, 0x5, 0x0, 0xD1, 0xEB, 0x69, 0x29, 0x38, 0xEB, 0x69, 0x69, 0x38, 0x6B, 0x5, 0x0, 0x91, 0xEB,
		0x69, 0x29, 0x38, 0xEB, 0x69, 0x69, 0x38, 0x2B, 0xFE, 0xFF, 0x35, 0x0, 0x0, 0x80, 0xD2, 0xE1, 0x3, 0xF, 0xAA,
		0x21, 0x0, 0x9, 0x8B, 0x22, 0x0, 0x80, 0xD2, 0x70, 0x0, 0x80, 0xD2, 0x1, 0x10, 0x0, 0xD4, 0x20, 0x0, 0x80, 0xD2,
		0xE1, 0x3, 0xF, 0xAA, 0x21, 0x0, 0x9, 0x8B, 0x22, 0x0, 0x80, 0xD2, 0x90, 0x0, 0x80, 0xD2, 0x1, 0x10, 0x0, 0xD4,
		0xE0, 0x3, 0xF, 0xAA, 0xC0, 0x3, 0x5F, 0xD6,
	}) {
		t.Errorf("jit code does not match expected code")
	}
}
