package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	optypes "github.com/DonnieTD/Gorth/OpTypes"
	opcreators "github.com/DonnieTD/Gorth/Opcreators"
	utils "github.com/DonnieTD/Gorth/Utils"
)

func LoadProgram(path string) []utils.Tuple {
	if optypes.COUNT_OPS != 4 {
		panic("Update CURRENT_OPCOUNT LoadProgram")
	}

	readFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	var ops []utils.Tuple
	for _, line := range fileLines {
		tokens := strings.Split(line, " ")
		for _, token := range tokens {
			switch token {
			case ".":
				ops = append(ops, opcreators.Dump())
			case "+":
				ops = append(ops, opcreators.Plus())
			case "-":
				ops = append(ops, opcreators.Minus())
			default:
				tokenInt, _ := strconv.Atoi(token)
				ops = append(ops, opcreators.Push(tokenInt))

			}
		}
	}

	return ops
}
