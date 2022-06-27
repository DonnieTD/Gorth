package main

import (
	"fmt"
	"os"

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

	if subcommand == "sim" {
		if len(os.Args) < 3 {
			fmt.Println("")
			fmt.Println("ERROR: no path was provided to sim")
			fmt.Println("")
			os.Exit(1)
		}
		Program := LoadProgram(os.Args[2])
		SimulateProgram(Program)
	} else if subcommand == "com" {
		if len(os.Args) < 3 {
			fmt.Println("")
			fmt.Println("ERROR: no path was provided to com")
			fmt.Println("")
			os.Exit(1)
		}
		Program := LoadProgram(os.Args[2])
		CompileProgram(Program, "output.asm")
		utils.RunCMD("nasm -felf64 output.asm")
		utils.RunCMD("ld -o output output.o")
	} else {
		fmt.Printf("EROOR: unknown subcommand %v \n", subcommand)
		os.Exit(1)
	}

}
