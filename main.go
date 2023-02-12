package main

import (
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Printf("gobf: try '%s input.b'\n", os.Args[0])
		os.Exit(2)
	}

	inputData := parseInputData(os.Args[1])

	parser := Parser{}
	instructions, err := parser.Parse(inputData)

	if err != nil {
		log.Printf("unrecoverable parser error: %s\n", err)
		os.Exit(1)
	}

	var memorySize uint = 30_000
	memory := make([]uint8, memorySize)

	cpu := Cpu{
		memory: memory,
		code:   instructions,
		input:  io.Reader(os.Stdin),
		output: io.Writer(os.Stdout),
	}

	for cpu.HasInstructionsLeft() {
		cpu.Step()
	}
}

func parseInputData(arg string) string {
	if arg == "-" {
		stdin, err := io.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}

		return string(stdin)
	}

	contents, err := os.ReadFile(arg)
	if err != nil {
		panic(err)
	}

	return string(contents)
}
