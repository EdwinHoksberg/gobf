package instructions

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

	// The following instructions are unofficial and only used when the instruction optimizer is enabled.

	// Clear is an instruction for optimizing zero-ing out a register: [-]
	Clear
)

type Instruction struct {
	Name  InstructionType
	Value int
}

func (instruction *InstructionType) ToString() string {
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
	case Clear:
		return "Clear"
	case Unknown:
		return "Unknown"
	}

	return "Unknown"
}

func (instruction *Instruction) IsJump() bool {
	return instruction.Name == JumpIfZero || instruction.Name == JumpUnlessZero
}

func (instruction *Instruction) CanBeOptimized() bool {
	return instruction.Name == MoveRight ||
		instruction.Name == MoveLeft ||
		instruction.Name == Increment ||
		instruction.Name == Decrement
}
