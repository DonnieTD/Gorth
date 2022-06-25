package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	utils "github.com/DonnieTD/Gorth/Utils"
)

func CompileProgram(program []utils.Tuple, programName string) {
	if COUNT_OPS != 4 {
		fmt.Println("Update CURRENT_OPCOUNT CompileProgram")
		return
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
	datawriter.WriteString("global _start" + "\n")
	datawriter.WriteString("_start:" + "\n")

	for _, operation := range program {
		switch operation.Optype {
		case OP_PUSH:
			datawriter.WriteString(fmt.Sprintf("    ;;-- push %d --", operation.Parameters) + "\n")
			datawriter.WriteString(fmt.Sprintf("    push %d", operation.Parameters) + "\n")
		case OP_PLUS:
			datawriter.WriteString(fmt.Sprintf("    ;;-- plus %d --", operation.Parameters) + "\n")
			datawriter.WriteString(fmt.Sprintf("    ;;-- TODO: not implemented -- %d", operation.Parameters) + "\n")
		case OP_MINUS:
			datawriter.WriteString(fmt.Sprintf("    ;;-- minus %d --", operation.Parameters) + "\n")
			datawriter.WriteString(fmt.Sprintf("    ;;-- TODO: not implemented -- %d", operation.Parameters) + "\n")
		case OP_DUMP:
			datawriter.WriteString(fmt.Sprintf("    ;;-- dump %d --", operation.Parameters) + "\n")
			datawriter.WriteString(fmt.Sprintf("    ;;-- TODO: not implemented -- %d", operation.Parameters) + "\n")
		}
		datawriter.WriteString("    mov" + "\n")
	}

	datawriter.WriteString("    mov rax, 60" + "\n")
	datawriter.WriteString("    mov rdi, 0" + "\n")
	datawriter.WriteString("    syscall" + "\n")

	datawriter.Flush()
	file.Close()
}
