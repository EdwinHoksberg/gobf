package main

type Jit struct {
	memorySize uint
	code       []byte
	codeBlocks []CodeBlock
}

type CodeBlock struct {
	instruction Instruction
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
