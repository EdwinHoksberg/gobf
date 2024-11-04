package main

type Jit struct {
	code       []byte
	codeBlocks []CodeBlock
}

type CodeBlock struct {
	instruction Instruction
	offset      int
	link        *CodeBlock
}
