package main

import (
	"fmt"
	"os"

	lexer "github.com/DonnieTD/Gorth/Lexer"
	utils "github.com/DonnieTD/Gorth/Utils"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("")
		fmt.Println("ERROR: no subcommand was provided")
		fmt.Println("")
		utils.Usage(os.Args[0])
		os.Exit(1)
	}

	subcommand := os.Args[1]

	if len(os.Args) < 3 {
		fmt.Println("")
		fmt.Printf("ERROR: no path was provided to %v \n", subcommand)
		fmt.Println("")
		os.Exit(1)
	}

	NAHI := NAH{
		LEXER: lexer.New(os.Args[2]),
	}

	NAHI.LEXER.LoadProgram()
	Program := NAHI.LEXER.Lex()

	if subcommand == "sim" {
		SimulateProgram(Program)
	} else if subcommand == "com" {
		CompileProgram(Program, "output.asm")
	} else {
		fmt.Printf("EROOR: unknown subcommand %v \n", subcommand)
		os.Exit(1)
	}

}
