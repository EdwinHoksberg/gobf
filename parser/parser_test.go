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

			if err != nil {
				t.Errorf("expected no error, got %s", err)
			}

			if len(instructions) != 1 {
				t.Errorf("expected 1 instruction to be parsed, got %d", len(instructions))
			}

			if instructions[0].Name != test.instructionType {
				t.Errorf("got %s, want %s", instructions[0].Name.ToString(), test.instructionType.ToString())
			}
		})
	}
}

func TestParser_ParseUnknown(t *testing.T) {
	parser := NewParser()
	instructions, err := parser.Parse("x")

	if err != nil {
		t.Errorf("expected no error, got %s", err)
	}

	if len(instructions) != 0 {
		t.Errorf("expected 0 instructions to be parsed, got %d", len(instructions))
	}
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

			//for i, instruction := range instructions {
			//	if instruction.Name != test.instructions[i].Name {
			//		t.Errorf("got name %s, want name %s", test.instructions[i].Name.ToString(), instruction.Name.ToString())
			//	}
			//	if instruction.Value != test.instructions[i].Value {
			//		t.Errorf("got link %d, want link %d", test.instructions[i].Value, instruction.Value)
			//	}
			//}
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

			if err.Error() != fmt.Sprintf("no matching '%s' found", test.expectedMissing) {
				t.Errorf("parsing missing instruction did not error")
			}

			if len(instructions) != 0 {
				t.Errorf("parsing produced unexpected instructions")
			}
		})
	}
}

func FuzzParser_Parse(f *testing.F) {
	parser := NewParser()
	f.Fuzz(func(t *testing.T, input string) {
		parser.Parse(input)
	})
}
