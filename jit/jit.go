package jit

import "gobf/instructions"

type Jit struct {
	memorySize uint
	code       []byte
	codeBlocks []CodeBlock
}

type CodeBlock struct {
	instruction instructions.Instruction
	offset      int
	link        *CodeBlock
}

func NewJit(memorySize uint) *Jit {
	return &Jit{
		memorySize,
		make([]byte, 0),
		make([]CodeBlock, 0),
	}
}

func (jit *Jit) GeneratedCode() []byte {
	return jit.code
}
