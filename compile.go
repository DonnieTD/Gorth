package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	optypes "github.com/DonnieTD/Gorth/OpTypes"
	utils "github.com/DonnieTD/Gorth/Utils"
)

func CompileProgram(program []utils.Tuple, programName string) {
	if optypes.COUNT_OPS != 4 {
		panic("Update CURRENT_OPCOUNT CompileProgram")
	}

	if _, err := os.Stat("./" + programName); err == nil {
		e := os.Remove(programName)
		if e != nil {
			log.Fatal(e)
		}
	}

	file, err := os.OpenFile(programName, os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	datawriter := bufio.NewWriter(file)

	datawriter.WriteString("segment .text" + "\n")

	datawriter.WriteString("dump:\n")
	datawriter.WriteString("    mov     r9, -3689348814741910323\n")
	datawriter.WriteString("    sub     rsp, 40\n")
	datawriter.WriteString("    mov     BYTE [rsp+31], 10\n")
	datawriter.WriteString("    lea     rcx, [rsp+30]\n")
	datawriter.WriteString(".L2:\n")
	datawriter.WriteString("    mov     rax, rdi\n")
	datawriter.WriteString("    lea     r8, [rsp+32]\n")
	datawriter.WriteString("    mul     r9\n")
	datawriter.WriteString("    mov     rax, rdi\n")
	datawriter.WriteString("    sub     r8, rcx\n")
	datawriter.WriteString("    shr     rdx, 3\n")
	datawriter.WriteString("    lea     rsi, [rdx+rdx*4]\n")
	datawriter.WriteString("    add     rsi, rsi\n")
	datawriter.WriteString("    sub     rax, rsi\n")
	datawriter.WriteString("    add     eax, 48\n")
	datawriter.WriteString("    mov     BYTE [rcx], al\n")
	datawriter.WriteString("    mov     rax, rdi\n")
	datawriter.WriteString("    mov     rdi, rdx\n")
	datawriter.WriteString("    mov     rdx, rcx\n")
	datawriter.WriteString("    sub     rcx, 1\n")
	datawriter.WriteString("    cmp     rax, 9\n")
	datawriter.WriteString("    ja      .L2\n")
	datawriter.WriteString("    lea     rax, [rsp+32]\n")
	datawriter.WriteString("    mov     edi, 1\n")
	datawriter.WriteString("    sub     rdx, rax\n")
	datawriter.WriteString("    xor     eax, eax\n")
	datawriter.WriteString("    lea     rsi, [rsp+32+rdx]\n")
	datawriter.WriteString("    mov     rdx, r8\n")
	datawriter.WriteString("    mov     rax, 1\n")
	datawriter.WriteString("    syscall\n")
	datawriter.WriteString("    add     rsp, 40\n")
	datawriter.WriteString("    ret\n")

	datawriter.WriteString("global _start" + "\n")
	datawriter.WriteString("_start:" + "\n")

	for _, operation := range program {
		switch operation.Optype {
		case optypes.OP_PUSH:
			datawriter.WriteString(fmt.Sprintf("    ;;-- push %d --", operation.Parameters) + "\n")
			datawriter.WriteString(fmt.Sprintf("    push %d", operation.Parameters) + "\n")
		case optypes.OP_PLUS:
			datawriter.WriteString("    ;;-- plus %d -- \n")
			datawriter.WriteString("    pop rax \n")
			datawriter.WriteString("    pop rbx \n")
			datawriter.WriteString("    add rax, rbx \n")
			datawriter.WriteString("    push rax \n")

		case optypes.OP_MINUS:
			datawriter.WriteString("    ;;-- minus %d -- \n")
			datawriter.WriteString("    pop rax \n")
			datawriter.WriteString("    pop rbx \n")
			datawriter.WriteString("    sub rbx, rax \n")
			datawriter.WriteString("    push rbx \n")
		case optypes.OP_DUMP:
			datawriter.WriteString("    ;;-- dump %d -- \n")
			datawriter.WriteString("    pop rdi \n")
			datawriter.WriteString("    call dump\n")
		}
	}

	datawriter.WriteString("    mov rax, 60" + "\n")
	datawriter.WriteString("    mov rdi, 0" + "\n")
	datawriter.WriteString("    syscall" + "\n")

	datawriter.Flush()
	file.Close()
}
