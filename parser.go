package main

import (
	"errors"
)

type Parser struct{}

func (parser *Parser) Parse(input string) ([]Instruction, error) {
	var depth = 0
	var depthMap = map[int]int{}
	var counter = 0

	var validInstructions = 0
	for _, character := range input {
		instructionName := parser.instructionTypeFromCharacter(character)

		if instructionName != Unknown {
			validInstructions++
		}
	}

	instructions := make([]Instruction, 0, validInstructions)

	for _, character := range input {
		instructionName := parser.instructionTypeFromCharacter(character)

		if instructionName == Unknown {
			continue
		}

		if instructionName == JumpIfZero {
			depth++
			depthMap[depth] = counter
		}

		var jumpPoint = 0
		if instructionName == JumpUnlessZero {
			if depth == 0 {
				return nil, errors.New("no matching '[' found")
			}

			jumpPoint = depthMap[depth]
			delete(depthMap, depth)

			depth--
			instructions[jumpPoint].jumpPoint = counter
		}

		instruction := Instruction{instructionName, jumpPoint}

		instructions = append(instructions, instruction)
		counter++
	}

	if len(depthMap) != 0 {
		return nil, errors.New("no matching ']' found")
	}

	return instructions, nil
}

func (parser *Parser) instructionTypeFromCharacter(character rune) InstructionType {
	switch character {
	case '>':
		return MoveRight
	case '<':
		return MoveLeft
	case '+':
		return Increment
	case '-':
		return Decrement
	case '.':
		return Write
	case ',':
		return Read
	case '[':
		return JumpIfZero
	case ']':
		return JumpUnlessZero
	default:
		return Unknown
	}
}
