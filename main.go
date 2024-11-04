package main

import (
	"flag"
	"golang.org/x/sys/unix"
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

	terminalSettings := updateTerminalSettings()
	defer resetTerminalSettings(terminalSettings)

	parser := Parser{}
	instructions, err := parser.Parse(inputData)

	if err != nil {
		log.Printf("unrecoverable parser error: %s\n", err)
		os.Exit(1)
	}

	jit := Jit{}
	if err := jit.compile(instructions); err != nil {
		panic(err)
	}

	if err := jit.run(*memorySize); err != nil {
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

func updateTerminalSettings() *unix.Termios {
	oldState, err := unix.IoctlGetTermios(int(os.Stdin.Fd()), getTermios)
	if err != nil {
		return nil
	}

	newState := oldState
	newState.Lflag &^= unix.ICANON | unix.ECHO

	if err := unix.IoctlSetTermios(int(os.Stdin.Fd()), setTermios, newState); err != nil {
		panic("failed to update terminal settings: " + err.Error())
	}

	return oldState
}

func resetTerminalSettings(terminalSettings *unix.Termios) {
	if terminalSettings == nil {
		return
	}

	if err := unix.IoctlSetTermios(int(os.Stdin.Fd()), setTermios, terminalSettings); err != nil {
		panic("failed to reset terminal settings: " + err.Error())
	}
}
