package parser

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gobf/instructions"
	"testing"
)

func TestParser_ParseSingle(t *testing.T) {
	var tests = []struct {
		input           string
		instructionType instructions.InstructionType
	}{
		{"+", instructions.Increment},
		{"-", instructions.Decrement},
		{">", instructions.MoveRight},
		{"<", instructions.MoveLeft},
		{",", instructions.Read},
		{".", instructions.Write},
	}

	for _, test := range tests {
		parser := NewParser()
		testName := fmt.Sprintf("%s == %s", test.input, test.instructionType.ToString())
		t.Run(testName, func(t *testing.T) {
			instructions, err := parser.Parse(test.input)

			assert.NoError(t, err)

			assert.Len(t, instructions, 1)

			assert.Equal(t, test.instructionType, instructions[0].Name)
		})
	}
}

func TestParser_ParseUnknown(t *testing.T) {
	parser := NewParser()
	instructions, err := parser.Parse("x")

	assert.NoError(t, err)

	assert.Empty(t, instructions)
}

func TestParser_ParseMultiple(t *testing.T) {
	var tests = []struct {
		input        string
		instructions []instructions.Instruction
	}{
		{
			"+-",
			[]instructions.Instruction{
				{
					Name:  instructions.Increment,
					Value: 1,
				},
				{
					Name:  instructions.Decrement,
					Value: 1,
				},
			},
		},
		{
			">><>",
			[]instructions.Instruction{
				{
					Name:  instructions.MoveRight,
					Value: 1,
				},
				{
					Name:  instructions.MoveRight,
					Value: 1,
				},
				{
					Name:  instructions.MoveLeft,
					Value: 1,
				},
				{
					Name:  instructions.MoveRight,
					Value: 1,
				},
			},
		},
		{
			"[[]][]",
			[]instructions.Instruction{
				{
					Name:  instructions.JumpIfZero,
					Value: 3,
				},
				{
					Name:  instructions.JumpIfZero,
					Value: 2,
				},
				{
					Name:  instructions.JumpUnlessZero,
					Value: 1,
				},
				{
					Name:  instructions.JumpUnlessZero,
					Value: 0,
				},
				{
					Name:  instructions.JumpIfZero,
					Value: 5,
				},
				{
					Name:  instructions.JumpUnlessZero,
					Value: 4,
				},
			},
		},
	}

	for _, test := range tests {
		parser := NewParser()
		t.Run(test.input, func(t *testing.T) {
			instructions, err := parser.Parse(test.input)

			assert.NoError(t, err)

			assert.Equal(t, test.instructions, instructions)
		})
	}
}

func TestParser_MatchJumpGroups(t *testing.T) {
	ifZero := instructions.JumpIfZero
	unlessZero := instructions.JumpUnlessZero

	var tests = []struct {
		name            string
		input           string
		expectedMissing string
	}{
		{ifZero.ToString(), "[", "]"},
		{unlessZero.ToString(), "]", "["},
	}

	for _, test := range tests {
		parser := NewParser()
		t.Run(test.name, func(t *testing.T) {
			instructions, err := parser.Parse(test.input)

			assert.Equal(t, fmt.Sprintf("no matching '%s' found", test.expectedMissing), err.Error())

			assert.Empty(t, instructions)
		})
	}
}

func FuzzParser_Parse(f *testing.F) {
	parser := NewParser()
	f.Fuzz(func(t *testing.T, input string) {
		parser.Parse(input)
	})
}
