package main

import (
	"flag"
	"io"
	"log"
	"os"
)

func main() {
	memorySize := flag.Uint("memory-size", 30_000, "Size (in bytes) of the memory available to the program")
	flag.Parse()

	if len(flag.Args()) != 1 {
		log.Printf("gobf: try '%s input.b'\n", os.Args[0])
		os.Exit(2)
	}

	inputData := parseInput(flag.Arg(0))

	terminalSettings := disableTerminalInputBuffering()
	defer resetTerminal(terminalSettings)

	parser := NewParser()
	instructions, err := parser.Parse(inputData)

	if err != nil {
		log.Printf("unrecoverable parser error: %s\n", err)
		os.Exit(1)
	}

	jit := NewJit()
	if err := jit.Compile(instructions); err != nil {
		panic(err)
	}

	if err := jit.Run(*memorySize); err != nil {
		panic(err)
	}
}

func parseInput(arg string) string {
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
