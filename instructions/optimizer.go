package instructions

import "fmt"

// @todo
//   - recalculate jump points
//   - actually use optimized value in jit code

type optimizedBlock struct {
	startIndex  int
	length      int
	instruction InstructionType
}

func OptimizeInstructions(instructions []Instruction) []Instruction {
	optimizedInstructions := make([]Instruction, 0)
	performedOptimizations := make([]optimizedBlock, 0)

	for instructionIndex := 0; instructionIndex < len(instructions); instructionIndex++ {
		instruction := instructions[instructionIndex]
		instructionType := instruction.Name

		// Only try to optimize instructions that can be optimized and if we still have instructions left to process
		if !instruction.CanBeOptimized() || len(instructions)-1 == instructionIndex {
			optimizedInstructions = append(optimizedInstructions, instruction)
			continue
		}

		// Check if we have at least one other instruction of the same type coming up
		if instructions[instructionIndex+1].Name != instructionType {
			optimizedInstructions = append(optimizedInstructions, instruction)
			continue
		}

		// Now we know what we have at least one more recurring instruction ahead,
		// loop until we find one that isn't the same, or we have reached the end of the instructions to process
		recurringInstructions := 1
		for j := 0; true; j++ {
			if instructionIndex+j > len(instructions)-1 || instructions[instructionIndex+j].Name != instructionType {
				recurringInstructions = j
				break
			}
		}

		fmt.Printf("optimizing %s, %d recurring\n", instructionType.ToString(), recurringInstructions)

		instruction.Value = recurringInstructions
		optimizedInstructions = append(optimizedInstructions, instruction)

		// Store optimizations that have been to done to patch up jumps later
		performedOptimizations = append(performedOptimizations, optimizedBlock{
			startIndex:  instructionIndex,
			length:      recurringInstructions - 1,
			instruction: instructionType,
		})

		instructionIndex += recurringInstructions - 1
	}

	recalculateJumps(&optimizedInstructions, performedOptimizations)

	fmt.Printf("before %d, after: %d (diff: %d)\n", len(instructions), len(optimizedInstructions), len(instructions)-len(optimizedInstructions))

	return optimizedInstructions
}

func recalculateJumps(instructions *[]Instruction, performedOptimizations []optimizedBlock) {
	recalculatedJumps := make(map[int]int)

	for _, optimization := range performedOptimizations {
		for instructionIndex := 0; instructionIndex < len(*instructions); instructionIndex++ {
			instruction := &(*instructions)[instructionIndex]

			if !instruction.IsJump() {
				continue
			}

			if _, ok := recalculatedJumps[instructionIndex]; !ok {
				recalculatedJumps[instructionIndex] = instruction.Value
			}

			// Fix jumps with links after offset and ...
			jumpLinksToInstructionAfterOffset := instruction.Value > optimization.startIndex && (optimization.startIndex+optimization.length) < instruction.Value
			// Fix jumps that exist after offset that link back to everything before offset+length
			jumpExistsAfterOffsetAndLinksBackBeforeOptimization := instructionIndex > optimization.startIndex && instruction.Value > optimization.startIndex

			if jumpLinksToInstructionAfterOffset || jumpExistsAfterOffsetAndLinksBackBeforeOptimization {
				recalculatedJumps[instructionIndex] -= optimization.length
			}
		}
	}

	for instructionIndex, newLink := range recalculatedJumps {
		(*instructions)[instructionIndex].Value = newLink
	}
}
