package instructions

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInstructions_Optimize(t *testing.T) {
	var tests = []struct {
		instructions         []Instruction
		expectedInstructions []Instruction
	}{
		{
			[]Instruction{
				{Name: Increment, Value: 1},
				{Name: Increment, Value: 1},
				{Name: Increment, Value: 1},
				{Name: Decrement, Value: 1},
				{Name: Decrement, Value: 1},
				{Name: Increment, Value: 1},
				{Name: MoveRight, Value: 1},
				{Name: MoveRight, Value: 1},
				{Name: MoveLeft, Value: 1},
				{Name: MoveLeft, Value: 1},
			},
			[]Instruction{
				{Name: Increment, Value: 3},
				{Name: Decrement, Value: 2},
				{Name: Increment, Value: 1},
				{Name: MoveRight, Value: 2},
				{Name: MoveLeft, Value: 2},
			},
		},
		{
			[]Instruction{
				{Name: Increment, Value: 1},
				{Name: JumpIfZero, Value: 4},
				{Name: Increment, Value: 1},
				{Name: Increment, Value: 1},
				{Name: JumpUnlessZero, Value: 1},
				{Name: Increment, Value: 1},
				{Name: JumpIfZero, Value: 7},
				{Name: JumpUnlessZero, Value: 6},
			},
			[]Instruction{
				{Name: Increment, Value: 1},
				{Name: JumpIfZero, Value: 3},
				{Name: Increment, Value: 2},
				{Name: JumpUnlessZero, Value: 1},
				{Name: Increment, Value: 1},
				{Name: JumpIfZero, Value: 6},
				{Name: JumpUnlessZero, Value: 5},
			},
		},
		{
			[]Instruction{
				{Name: Increment, Value: 0}, // 13
				{Name: Increment, Value: 0},
				{Name: Increment, Value: 0},
				{Name: Increment, Value: 0},
				{Name: Increment, Value: 0},
				{Name: Increment, Value: 0},
				{Name: Increment, Value: 0},
				{Name: Increment, Value: 0},
				{Name: Increment, Value: 0},
				{Name: Increment, Value: 0},
				{Name: Increment, Value: 0},
				{Name: Increment, Value: 0},
				{Name: Increment, Value: 0},

				{Name: JumpIfZero, Value: 37},

				{Name: Decrement, Value: 0},

				{Name: MoveRight, Value: 0},

				{Name: Increment, Value: 0}, // 2
				{Name: Increment, Value: 0},

				{Name: MoveRight, Value: 0}, // 3
				{Name: MoveRight, Value: 0},
				{Name: MoveRight, Value: 0},

				{Name: Increment, Value: 0}, // 5
				{Name: Increment, Value: 0},
				{Name: Increment, Value: 0},
				{Name: Increment, Value: 0},
				{Name: Increment, Value: 0},

				{Name: MoveRight, Value: 0},

				{Name: Increment, Value: 0}, // 2
				{Name: Increment, Value: 0},

				{Name: MoveRight, Value: 0},

				{Name: Increment, Value: 0},

				{Name: MoveLeft, Value: 0}, // 6
				{Name: MoveLeft, Value: 0},
				{Name: MoveLeft, Value: 0},
				{Name: MoveLeft, Value: 0},
				{Name: MoveLeft, Value: 0},
				{Name: MoveLeft, Value: 0},

				{Name: JumpUnlessZero, Value: 13},

				{Name: MoveRight, Value: 0}, // 5
				{Name: MoveRight, Value: 0},
				{Name: MoveRight, Value: 0},
				{Name: MoveRight, Value: 0},
				{Name: MoveRight, Value: 0},
			},
			[]Instruction{
				{Name: Increment, Value: 13}, // 13

				{Name: JumpIfZero, Value: 12},

				{Name: Decrement, Value: 0},

				{Name: MoveRight, Value: 0},

				{Name: Increment, Value: 2}, // 2

				{Name: MoveRight, Value: 3}, // 3

				{Name: Increment, Value: 5}, // 5

				{Name: MoveRight, Value: 0},

				{Name: Increment, Value: 2}, // 2

				{Name: MoveRight, Value: 0},

				{Name: Increment, Value: 0},

				{Name: MoveLeft, Value: 6}, // 6

				{Name: JumpUnlessZero, Value: 1},

				{Name: MoveRight, Value: 5}, // 5
			},
		},
		{
			[]Instruction{
				{Name: Increment, Value: 0},
				{Name: Increment, Value: 0},
				{Name: JumpIfZero, Value: 6},
				{Name: Increment, Value: 0},
				{Name: Increment, Value: 0},
				{Name: Decrement, Value: 0},
				{Name: Decrement, Value: 0},
				{Name: JumpUnlessZero, Value: 2},
				{Name: Increment, Value: 0},
			},
			[]Instruction{
				{Name: Increment, Value: 2},
				{Name: JumpIfZero, Value: 4},
				{Name: Increment, Value: 2},
				{Name: Decrement, Value: 2},
				{Name: JumpUnlessZero, Value: 1},
				{Name: Increment, Value: 0},
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test_%d", i+1), func(t *testing.T) {
			optimizedInstructions := OptimizeInstructions(test.instructions)

			assert.Equal(t, test.expectedInstructions, optimizedInstructions)
		})
	}
}
