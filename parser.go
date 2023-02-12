package main

import (
	"errors"
)

type Parser struct {
	depth    int
	depthMap map[int]int
}

func (parser *Parser) Parse(input string) ([]Instruction, error) {
	parser.depth = 0
	parser.depthMap = map[int]int{}

	var instructions []Instruction
	var counter int = 0

	for _, character := range input {
		instructionName := parser.instructionTypeFromCharacter(character)

		if instructionName == Unknown {
			continue
		}

		if instructionName == JumpIfZero {
			parser.depth++
			parser.depthMap[parser.depth] = counter
		}

		var jumpPoint = 0
		if instructionName == JumpUnlessZero {
			if parser.depth == 0 {
				return nil, errors.New("no matching '[' found")
			}

			jumpPoint = parser.depthMap[parser.depth]
			delete(parser.depthMap, parser.depth)

			parser.depth--
			instructions[jumpPoint].jumpPoint = counter
		}

		instruction := Instruction{instructionName, jumpPoint}

		instructions = append(instructions, instruction)
		counter++
	}

	if len(parser.depthMap) != 0 {
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
