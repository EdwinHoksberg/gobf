package main

type InstructionType uint

const (
	Unknown InstructionType = iota
	MoveRight
	MoveLeft
	Increment
	Decrement
	Write
	Read
	JumpIfZero
	JumpUnlessZero
)

type Instruction struct {
	name InstructionType
	link int
}

func (instruction *InstructionType) toString() string {
	switch *instruction {
	case MoveRight:
		return "MoveRight"
	case MoveLeft:
		return "MoveLeft"
	case Increment:
		return "Increment"
	case Decrement:
		return "Decrement"
	case Write:
		return "Write"
	case Read:
		return "Read"
	case JumpIfZero:
		return "JumpIfZero"
	case JumpUnlessZero:
		return "JumpUnlessZero"
	case Unknown:
		return "Unknown"
	}

	return "Unknown"
}
