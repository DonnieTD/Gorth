package main

import (
	"bufio"
	"log"
	"os"

	utils "github.com/DonnieTD/Gorth/Utils"
)

func CompileProgram(program []utils.Tuple, programName string) {
	file, err := os.OpenFile(programName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	datawriter := bufio.NewWriter(file)
	datawriter.WriteString("segment .text" + "\n")
	datawriter.WriteString("global _start" + "\n")
	datawriter.WriteString("_start:" + "\n")
	datawriter.WriteString("    mov rax, 60" + "\n")
	datawriter.WriteString("    mov rdi, 0" + "\n")
	datawriter.WriteString("    syscall" + "\n")

	datawriter.Flush()
	file.Close()
}
