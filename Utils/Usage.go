package utils

import "fmt"

func Usage(programName string) {
	fmt.Printf("Usage: %s <SUBCOMMAND> [ARGS]\n", programName)
	fmt.Println("")
	fmt.Println("SUBCOMMANDS:")
	fmt.Println("    interpret        Interpret the program")
	fmt.Println("    compile        Compile the program")
	fmt.Println("")
}
