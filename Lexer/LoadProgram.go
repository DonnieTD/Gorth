package lexer

import (
	"bufio"
	"fmt"
	"os"
)

func (lex *Lexer) LoadProgram() {
	readFile, err := os.Open(lex.FilePath)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines [][]rune

	for fileScanner.Scan() {
		fileLines = append(fileLines, []rune(fileScanner.Text()))
	}

	lex.Program = fileLines

	readFile.Close()
}
