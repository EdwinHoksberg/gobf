package main

import (
	"encoding/hex"
	"flag"
	"gobf/jit"
	"gobf/parser"
	"io"
	"log"
	"os"
)

func main() {
	memorySize := flag.Uint("memory-size", 30_000, "Size (in bytes) of the memory available to the program")
	dumpGeneratedJitCode := flag.Bool("dump-jit", false, "Dump generated JIT code to stderr")
	flag.Parse()

	if len(flag.Args()) != 1 {
		log.Printf("gobf: try '%s input.b'\n", os.Args[0])
		os.Exit(2)
	}

	inputData := parseInput(flag.Arg(0))

	terminalSettings := disableTerminalInputBuffering()
	defer resetTerminal(terminalSettings)

	instructionParser := parser.NewParser()
	instructions, err := instructionParser.Parse(inputData)

	if err != nil {
		log.Printf("unrecoverable parser error: %s\n", err)
		os.Exit(1)
	}

	jitter := jit.NewJit(*memorySize)
	if err := jitter.Compile(instructions); err != nil {
		panic(err)
	}

	if *dumpGeneratedJitCode {
		if _, err := os.Stderr.WriteString(hex.EncodeToString(jitter.GeneratedCode())); err != nil {
			log.Printf("error writing jit code to stderr: %s\n", err)
		}
	}

	if err := jitter.Run(); err != nil {
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
