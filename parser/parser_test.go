package parser

import (
	"fmt"
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
					Name: instructions.Increment,
					Link: 0,
				},
				{
					Name: instructions.Decrement,
					Link: 0,
				},
			},
		},
		{
			">><>",
			[]instructions.Instruction{
				{
					Name: instructions.MoveRight,
					Link: 0,
				},
				{
					Name: instructions.MoveRight,
					Link: 0,
				},
				{
					Name: instructions.MoveLeft,
					Link: 0,
				},
				{
					Name: instructions.MoveRight,
					Link: 0,
				},
			},
		},
		{
			"[[]][]",
			[]instructions.Instruction{
				{
					Name: instructions.JumpIfZero,
					Link: 3,
				},
				{
					Name: instructions.JumpIfZero,
					Link: 2,
				},
				{
					Name: instructions.JumpUnlessZero,
					Link: 1,
				},
				{
					Name: instructions.JumpUnlessZero,
					Link: 0,
				},
				{
					Name: instructions.JumpIfZero,
					Link: 5,
				},
				{
					Name: instructions.JumpUnlessZero,
					Link: 4,
				},
			},
		},
	}

	for _, test := range tests {
		parser := NewParser()
		t.Run(test.input, func(t *testing.T) {
			instructions, err := parser.Parse(test.input)

			if err != nil {
				t.Errorf("expected no error, got %s", err)
			}

			if len(instructions) != len(test.instructions) {
				t.Errorf("expected %d instruction to be parsed, got %d (%v)", len(test.instructions), len(instructions), instructions)
			}

			for i, instruction := range instructions {
				if instruction.Name != test.instructions[i].Name {
					t.Errorf("got name %s, want name %s", test.instructions[i].Name.ToString(), instruction.Name.ToString())
				}
				if instruction.Link != test.instructions[i].Link {
					t.Errorf("got link %d, want link %d", test.instructions[i].Link, instruction.Link)
				}
			}
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
