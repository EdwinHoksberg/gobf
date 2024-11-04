package main

import (
	"golang.org/x/sys/unix"
	"os"
)

func disableTerminalInputBuffering() *unix.Termios {
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

func resetTerminal(terminalSettings *unix.Termios) {
	if terminalSettings == nil {
		return
	}

	if err := unix.IoctlSetTermios(int(os.Stdin.Fd()), setTermios, terminalSettings); err != nil {
		panic("failed to reset terminal settings: " + err.Error())
	}
}
