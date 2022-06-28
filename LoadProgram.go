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

func TextTokenToTuple(token string, lineNumber string, tokenPosition string) utils.Tuple {
	if optypes.COUNT_OPS != 4 {
		panic("Update CURRENT_OPCOUNT LoadProgram")
	}

	switch token {
	case ".":
		return opcreators.Dump(lineNumber, tokenPosition)
	case "+":
		return opcreators.Plus(lineNumber, tokenPosition)
	case "-":
		return opcreators.Minus(lineNumber, tokenPosition)
	default:
		tokenInt, _ := strconv.Atoi(token)
		return opcreators.Push(tokenInt, lineNumber, tokenPosition)
	}
}

func LoadProgram(path string) []utils.Tuple {
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
	for lineNumber, line := range fileLines {
		chars := strings.Split(line, "")
		new_token_buffer := []string{}
		for currentTokenIndex, char := range chars {
			if char == " " {
				if len(new_token_buffer) == 0 {
					continue
				} else {
					new_token := strings.Join(new_token_buffer, "")
					// Place in tokens that are in the buffer when encountering a space
					// Everything will be indexed from 1 ( ??? ) is this a bad choice ( vs code lines arent zero indexed )
					TTokenTuple := TextTokenToTuple(new_token, fmt.Sprint(lineNumber+1), fmt.Sprint(currentTokenIndex-(len(new_token)-1)))
					ops = append(ops, TTokenTuple)
					new_token_buffer = []string{}
					continue
				}
			} else {
				new_token_buffer = append(new_token_buffer, char)

				//  place tokens in at end if it doesnt end ons space
				if currentTokenIndex == len(chars) {
					new_token := strings.Join(new_token_buffer, "")
					TTokenTuple := TextTokenToTuple(new_token, fmt.Sprint(lineNumber+1), fmt.Sprint(currentTokenIndex-(len(new_token)-1)))
					ops = append(ops, TTokenTuple)
					new_token_buffer = []string{}
					continue
				}
			}
		}
	}

	return ops
}
