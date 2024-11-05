package main

import (
	"fmt"
	"testing"
)

func TestParser_ParseSingle(t *testing.T) {
	var tests = []struct {
		input           string
		instructionType InstructionType
	}{
		{"+", Increment},
		{"-", Decrement},
		{">", MoveRight},
		{"<", MoveLeft},
		{",", Read},
		{".", Write},
	}

	for _, test := range tests {
		parser := NewParser()
		testName := fmt.Sprintf("%s == %s", test.input, test.instructionType.toString())
		t.Run(testName, func(t *testing.T) {
			instructions, err := parser.Parse(test.input)

			if err != nil {
				t.Errorf("expected no error, got %s", err)
			}

			if len(instructions) != 1 {
				t.Errorf("expected 1 instruction to be parsed, got %d", len(instructions))
			}

			if instructions[0].name != test.instructionType {
				t.Errorf("got %s, want %s", instructions[0].name.toString(), test.instructionType.toString())
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
		instructions []Instruction
	}{
		{
			"+-",
			[]Instruction{
				{
					name: Increment,
					link: 0,
				},
				{
					name: Decrement,
					link: 0,
				},
			},
		},
		{
			">><>",
			[]Instruction{
				{
					name: MoveRight,
					link: 0,
				},
				{
					name: MoveRight,
					link: 0,
				},
				{
					name: MoveLeft,
					link: 0,
				},
				{
					name: MoveRight,
					link: 0,
				},
			},
		},
		{
			"[[]][]",
			[]Instruction{
				{
					name: JumpIfZero,
					link: 3,
				},
				{
					name: JumpIfZero,
					link: 2,
				},
				{
					name: JumpUnlessZero,
					link: 1,
				},
				{
					name: JumpUnlessZero,
					link: 0,
				},
				{
					name: JumpIfZero,
					link: 5,
				},
				{
					name: JumpUnlessZero,
					link: 4,
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
				if instruction.name != test.instructions[i].name {
					t.Errorf("got name %s, want name %s", test.instructions[i].name.toString(), instruction.name.toString())
				}
				if instruction.link != test.instructions[i].link {
					t.Errorf("got link %d, want link %d", test.instructions[i].link, instruction.link)
				}
			}
		})
	}
}

func TestParser_MatchJumpGroups(t *testing.T) {
	ifZero := JumpIfZero
	unlessZero := JumpUnlessZero

	var tests = []struct {
		name            string
		input           string
		expectedMissing string
	}{
		{ifZero.toString(), "[", "]"},
		{unlessZero.toString(), "]", "["},
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
