package utils

import "fmt"

func Usage() {
	fmt.Println("Usage: gorth <SUBCOMMAND> [ARGS]")
	fmt.Println("")
	fmt.Println("SUBCOMMANDS:")
	fmt.Println("    sim        Simulate the program")
	fmt.Println("    com        Compile the program")
	fmt.Println("")
}
