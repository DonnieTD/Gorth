package main

import (
	"fmt"

	utils "github.com/DonnieTD/Gorth/Utils"
)

func CompileProgram(program []utils.Tuple,programName string) {
	file, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}


	fmt.Println("Simulate not implemented")
}
