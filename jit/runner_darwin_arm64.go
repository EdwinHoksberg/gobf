//go:build darwin && arm64

package jit

import (
	"errors"
	"syscall"
	"unsafe"
)

func (jit *Jit) Run() error {
	// Allocate program memory
	programMemory, err := syscall.Mmap(-1, 0, int(jit.memorySize), syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_ANON|syscall.MAP_PRIVATE)
	if err != nil {
		return errors.New("failed to map program memory: " + err.Error())
	}

	// Allocate executable memory
	executableMemory, err := syscall.Mmap(-1, 0, len(jit.code), syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_ANON|syscall.MAP_PRIVATE)
	if err != nil {
		return errors.New("failed to map executable memory: " + err.Error())
	}

	// Copy JIT instructions to executable memory
	copy(executableMemory, jit.code)

	// Change permissions to executable
	if err := syscall.Mprotect(executableMemory, syscall.PROT_EXEC); err != nil {
		return errors.New("failed to make executable memory executable: " + err.Error())
	}

	// Cast allocated memory regions to function pointers
	executableMemoryPointer := &executableMemory
	programMemoryPointer := unsafe.Pointer(&programMemory[0])

	// Define JIT call function and execute it
	f := *(*func(programMemory unsafe.Pointer))(unsafe.Pointer(&executableMemoryPointer))
	f(programMemoryPointer)

	return nil
}
