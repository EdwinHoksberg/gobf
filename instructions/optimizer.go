package instructions

type optimizedBlock struct {
	startIndex int
	length     int
}

var optimizedInstructions []Instruction
var performedOptimizations []optimizedBlock

func OptimizeInstructions(instructions []Instruction) []Instruction {
	optimizedInstructions = make([]Instruction, 0)
	performedOptimizations = make([]optimizedBlock, 0)

	for instructionIndex := 0; instructionIndex < len(instructions); instructionIndex++ {
		instruction := instructions[instructionIndex]
		instructionType := instruction.Name

		// Only try to optimize if we still have instructions left to process
		if len(instructions)-1 == instructionIndex {
			optimizedInstructions = append(optimizedInstructions, instruction)
			continue
		}

		// Clear instruction optimization (needs at least 3 instructions)
		if instructionIndex < len(instructions)-2 {
			// [-]
			if instructionType == JumpIfZero &&
				instructions[instructionIndex+1].Name == Decrement &&
				instructions[instructionIndex+2].Name == JumpUnlessZero {

				instruction.Name = Clear

				storeOptimization(&instructionIndex, 2, instruction)
				continue
			}
		}

		// Only try to optimize instructions that can be optimized
		if !instruction.CanBeOptimized() {
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

		instruction.Value = recurringInstructions
		storeOptimization(&instructionIndex, recurringInstructions-1, instruction)
	}

	recalculateJumps(&optimizedInstructions, performedOptimizations)

	return optimizedInstructions
}

func storeOptimization(instructionIndex *int, length int, instruction Instruction) {
	optimizedInstructions = append(optimizedInstructions, instruction)

	// Store optimizations that have been to done to patch up jumps later
	performedOptimizations = append(performedOptimizations, optimizedBlock{
		startIndex: *instructionIndex,
		length:     length,
	})

	*instructionIndex += length
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

			// Fix jumps with links after offset
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
