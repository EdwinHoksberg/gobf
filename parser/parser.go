package parser

import (
	"errors"
	"gobf/instructions"
)

type Parser struct{}

func NewParser() Parser {
	return Parser{}
}

func (parser *Parser) Parse(input string) ([]instructions.Instruction, error) {
	var depth = 0
	var depthMap = map[int]int{}
	var counter = 0

	var validInstructions = 0
	for _, character := range input {
		instructionName := parser.instructionTypeFromCharacter(character)

		if instructionName != instructions.Unknown {
			validInstructions++
		}
	}

	parsedInstructions := make([]instructions.Instruction, 0, validInstructions)

	for _, character := range input {
		instructionName := parser.instructionTypeFromCharacter(character)

		if instructionName == instructions.Unknown {
			continue
		}

		instructionValue := 1
		if instructionName == instructions.JumpIfZero {
			depth++
			depthMap[depth] = counter

			instructionValue = 0
		}

		if instructionName == instructions.JumpUnlessZero {
			if depth == 0 {
				return nil, errors.New("no matching '[' found")
			}

			instructionValue = depthMap[depth]
			delete(depthMap, depth)

			depth--
			parsedInstructions[instructionValue].Value = counter
		}

		instruction := instructions.Instruction{
			Name:  instructionName,
			Value: instructionValue,
		}

		parsedInstructions = append(parsedInstructions, instruction)
		counter++
	}

	if len(depthMap) != 0 {
		return nil, errors.New("no matching ']' found")
	}

	return parsedInstructions, nil
}

func (parser *Parser) instructionTypeFromCharacter(character rune) instructions.InstructionType {
	switch character {
	case '>':
		return instructions.MoveRight
	case '<':
		return instructions.MoveLeft
	case '+':
		return instructions.Increment
	case '-':
		return instructions.Decrement
	case '.':
		return instructions.Write
	case ',':
		return instructions.Read
	case '[':
		return instructions.JumpIfZero
	case ']':
		return instructions.JumpUnlessZero
	default:
		return instructions.Unknown
	}
}
