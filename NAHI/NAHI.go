package nahi

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"

	lexer "github.com/DonnieTD/NAH/Lexer"
	utils "github.com/DonnieTD/NAH/Utils"
)

type NAH struct {
	LEXER *lexer.Lexer
}

func GenerateAssemblyForDump(datawriter *bufio.Writer) {
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
}

func (n *NAH) Compile() {
	if lexer.COUNT_TOKENS != 4 {
		fmt.Println("Update CURRENT_OPCOUNT CompileProgram")
		os.Exit(1)
	}

	if _, err := os.Stat("./" + "output.asm"); err == nil {
		e := os.Remove("output.asm")
		if e != nil {
			log.Fatal(e)
		}
	}

	file, err := os.OpenFile("output.asm", os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	datawriter := bufio.NewWriter(file)
	datawriter.WriteString("segment .text" + "\n")
	GenerateAssemblyForDump(datawriter)
	datawriter.WriteString("global _start" + "\n")
	datawriter.WriteString("_start:" + "\n")

	for _, token := range n.LEXER.Tokens {
		switch token.TokenType {
		case lexer.TOKEN_PUSH:
			datawriter.WriteString(fmt.Sprintf("    ;;-- push %d --", token.Parameter) + "\n")
			datawriter.WriteString(fmt.Sprintf("    push %d", token.Parameter) + "\n")
		case lexer.TOKEN_PLUS:
			datawriter.WriteString("    ;;-- plus %d -- \n")
			datawriter.WriteString("    pop rax \n")
			datawriter.WriteString("    pop rbx \n")
			datawriter.WriteString("    add rax, rbx \n")
			datawriter.WriteString("    push rax \n")
		case lexer.TOKEN_MINUS:
			datawriter.WriteString("    ;;-- minus %d -- \n")
			datawriter.WriteString("    pop rax \n")
			datawriter.WriteString("    pop rbx \n")
			datawriter.WriteString("    sub rbx, rax \n")
			datawriter.WriteString("    push rbx \n")
		case lexer.TOKEN_DUMP:
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

	utils.RunCMD("nasm -felf64 output.asm")
	utils.RunCMD("ld -o output output.o")
}

func (n *NAH) Interpret() {
	if lexer.COUNT_TOKENS != 4 {
		fmt.Println("Update CURRENT_OPCOUNT SimulateProgram")
		os.Exit(1)
	}

	var programstack utils.Stack

	for _, token := range n.LEXER.Tokens {
		switch token.TokenType {
		case lexer.TOKEN_PUSH:
			programstack.Push(token.Parameter)
		case lexer.TOKEN_PLUS:
			a, _ := programstack.Pop()
			b, _ := programstack.Pop()
			if reflect.TypeOf(a).Kind() == reflect.Int && reflect.TypeOf(b).Kind() == reflect.Int {
				a := a.(int)
				b := b.(int)
				programstack.Push(a + b)
			}
			// later on do string concat here maybe
		case lexer.TOKEN_MINUS:
			a, _ := programstack.Pop()
			b, _ := programstack.Pop()
			if reflect.TypeOf(a).Kind() == reflect.Int && reflect.TypeOf(b).Kind() == reflect.Int {
				a := a.(int)
				b := b.(int)
				programstack.Push(b - a)
			}
		case lexer.TOKEN_DUMP:
			a, _ := programstack.Pop()
			fmt.Printf("%v \n", a)
		default:
			fmt.Println("Unreachable")
		}
	}
}
