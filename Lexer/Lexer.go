package lexer

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"unicode"
)

type Lexer struct {
	Cursor     int
	LineNumber int
	FilePath   string
	Program    [][]rune
	Tokens     []Token
}

type Token struct {
	TokenType  int
	Parameter  interface{}
	LineNumber int
	Position   int
}

func New(filePath string) *Lexer {
	lex := Lexer{
		// LINE NUMBERS AND COLS ARE 1 INDEXED IN REPORTING REMEMBER TO INCREMENT
		Cursor:     0,
		LineNumber: 0,
		FilePath:   filePath,
		Program:    [][]rune{},
		Tokens:     []Token{},
	}

	lex.LoadProgram()

	return &lex
}

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

func (lex *Lexer) TextToToken(text string) Token {
	if COUNT_TOKENS != 5 {
		abs, err := filepath.Abs("./Lexer/Lexer.go")
		if err == nil {
			fmt.Printf("Error in: %v\nUpdate CURRENT_OPCOUNT TextToToken() \n", abs)
		}
		os.Exit(1)
	}
	switch text {
	case ".":
		return Token{
			Position:   lex.Cursor,
			LineNumber: lex.LineNumber,
			TokenType:  TOKEN_DUMP,
			Parameter:  nil,
		}
	case "+":
		return Token{
			Position:   lex.Cursor,
			LineNumber: lex.LineNumber,
			TokenType:  TOKEN_PLUS,
			Parameter:  nil,
		}
	case "-":
		return Token{
			Position:   lex.Cursor,
			LineNumber: lex.LineNumber,
			TokenType:  TOKEN_MINUS,
			Parameter:  nil,
		}
	case "=":
		return Token{
			Position:   lex.Cursor,
			LineNumber: lex.LineNumber,
			TokenType:  TOKEN_EQUALS,
			Parameter:  nil,
		}
	default:
		// HERE WE MUST TEST IF THIS IS WORTHY OF A NUMBER CONVERSION BEFORE DOING SO
		tokenInt, err := strconv.Atoi(text)
		if err != nil {
			fmt.Printf("Error: Invalid NUMBER at %v:%v \n", lex.LineNumber+1, lex.Cursor-(len(text)-1)+1)
			os.Exit(1)
		}
		return Token{
			Position:   lex.Cursor - len(text),
			LineNumber: lex.LineNumber,
			TokenType:  TOKEN_PUSH,
			Parameter:  tokenInt,
		}
	}
}

func (lex *Lexer) LexLine(text []rune) {
	rune_buffer := []rune{}

	for index, char := range text {
		lex.Cursor = index
		if unicode.IsSpace(char) {
			lex.Cursor = index
			if len(rune_buffer) > 0 {
				lex.Tokens = append(lex.Tokens, lex.TextToToken(string(rune_buffer)))
				rune_buffer = []rune{}
				continue
			} else {
				continue
			}
		} else {
			rune_buffer = append(rune_buffer, char)
			if index == len(text)-1 {
				lex.Tokens = append(lex.Tokens, lex.TextToToken(string(rune_buffer)))
				rune_buffer = []rune{}
			}
			continue
		}
	}
}

func (lex *Lexer) Lex() {
	for index, line := range lex.Program {
		lex.LineNumber = index
		lex.LexLine(line)
	}
}
