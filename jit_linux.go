//go:build linux && amd64

package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

type Jit struct {
	output []byte
}

func (jit *Jit) compile(instructions []Instruction) error {
	//jit.code = append(jit.code,
	//	0x50,       // push rax
	//	0x53,       // push rbx
	//	0x41, 0x50, // push r8
	//	0x48, 0x31, 0xC0, // xor rax, rax
	//	0x48, 0x31, 0xDB, // xor rbx, rbx
	//	0x49, 0xC7, 0xC0, 0x01, 0x00, 0x00, 0x00, // mov r8, 1
	//)

	// program counter = rax
	// address counter = rbx
	// general purpose 1 = r8
	// general purpose 2 = r9

	//for _, instruction := range instructions {
	//	switch instruction.name {
	//	case MoveRight:
	//		jit.code = append(jit.code, 0x48, 0x83, 0xC3, 0x01) // add rbx, 1
	//	case MoveLeft:
	//		jit.code = append(jit.code, 0x48, 0x83, 0xEB, 0x01) // sub rbx, 1
	//	case Increment:
	//		jit.code = append(jit.code, 0xC6, 0x04, 0x1F, 0x01) // mov byte ptr [rdi + rbx], 1
	//	case Decrement:
	//		//
	//	case Write:
	//		//
	//	case Read:
	//		//
	//	case JumpIfZero:
	//		//
	//	case JumpUnlessZero:
	//		//
	//	}
	//}

	jit.output = append(jit.output,
		0xe0, 0x07, 0xbe, 0xa9, // stp x0, x1, [sp, #-32]!
		0xe2, 0x0b, 0x00, 0xf9, // str x2, [sp, #16]

		0x80, 0x00, 0x80, 0xd2, // mov x0, #0x4
		0x81, 0x00, 0x80, 0xd2, // mov x1, #0x4
		0x22, 0x7c, 0x00, 0x9b, // mul x2, x1, x0
		0x00, 0x00, 0x20, 0xd4, // brk

		0xe2, 0x0b, 0x40, 0xf9, // ldr x2, [sp, #16]
		0xe0, 0x07, 0xc2, 0xa8, // ldp x0, x1, [sp], #32

		0xc0, 0x03, 0x5f, 0xd6, // ret

		//0xcc,       // trap
		//0x41, 0x58, // pop r8
		//0x5B, // pop rbx
		//0x58, // pop rax
		//0xc3, // ret
	)

	return nil
}

func (jit *Jit) run() {
	//os.Stdout.Write(jit.code)

	// @todo mmap also the cpu.memory region for jit to use
	programMemory, err := syscall.Mmap(-1, 0, 64, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_ANON|syscall.MAP_PRIVATE)
	if err != nil {
		panic(err)
	}

	executableMemory, err := syscall.Mmap(-1, 0, len(jit.output), syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_ANON|syscall.MAP_PRIVATE)
	if err != nil {
		panic(err)
	}

	copy(executableMemory, jit.output)

	if err := syscall.Mprotect(executableMemory, syscall.PROT_READ|syscall.PROT_EXEC); err != nil {
		panic(err)
	}

	executableMemoryPointer := &executableMemory
	f := *(*func(m *[]byte))(unsafe.Pointer(&executableMemoryPointer))
	f(&programMemory)

	fmt.Printf("%x", programMemory)

	//print("x")
}
