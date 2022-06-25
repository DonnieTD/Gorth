package main

import (
	"fmt"
	"os"

	utils "github.com/DonnieTD/Gorth/Utils"
)

// TODO: unhardcode program
var Program = []utils.Tuple{
	Push(34),
	Push(35),
	Plus(),
	Dump(),
	Push(500),
	Push(80),
	Minus(),
	Dump(),
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("")
		fmt.Println("ERROR: no subcommand was provided")
		fmt.Println("")
		utils.Usage()
		os.Exit(1)
	}

	subcommand := os.Args[1]

	if subcommand == "sim" {
		SimulateProgram(Program)
	} else if subcommand == "com" {
		CompileProgram(Program, "output.asm")
		utils.RunCMD("nasm", "-felf64", "output.asm", "", "")
		utils.RunCMD("ld", "-0", "output", "output.o", "")

	} else {
		fmt.Printf("EROOR: unknown subcommand %v \n", subcommand)
		os.Exit(1)
	}

}
