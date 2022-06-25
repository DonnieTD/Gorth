package utils

import "fmt"

func Usage(programName string) {
	fmt.Printf("Usage: %s <SUBCOMMAND> [ARGS]\n",programName)
	fmt.Println("")
	fmt.Println("SUBCOMMANDS:")
	fmt.Println("    sim        Simulate the program")
	fmt.Println("    com        Compile the program")
	fmt.Println("")
}
