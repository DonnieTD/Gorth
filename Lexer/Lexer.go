package lexer

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type Lexer struct {
	cursor     int
	lineNumber int
	filePath   string
	program    [][]rune
	tokens     []Token
}

type Token struct {
	TokenType  int
	Parameter  interface{}
	LineNumber int
	Position   int
}

func New(filePath string) *Lexer {
	lex := &Lexer{
		// LINE NUMBERS AND COLS ARE 1 INDEXED IN REPORTING REMEMBER TO INCREMENT
		cursor:     0,
		lineNumber: 0,
		filePath:   filePath,
		program:    [][]rune{},
		tokens:     []Token{},
	}

	lex.LoadProgram()

	return lex
}

func (lex *Lexer) LoadProgram() {
	readFile, err := os.Open(lex.filePath)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines [][]rune

	for fileScanner.Scan() {
		fileLines = append(fileLines, []rune(fileScanner.Text()))
	}

	lex.program = fileLines

	readFile.Close()
}

func (lex *Lexer) TextToToken(text string) Token {
	switch text {
	case ".":
		return Token{
			Position:   lex.cursor,
			LineNumber: lex.lineNumber,
			TokenType:  TOKEN_DUMP,
			Parameter:  nil,
		}
	case "+":
		return Token{
			Position:   lex.cursor,
			LineNumber: lex.lineNumber,
			TokenType:  TOKEN_PLUS,
			Parameter:  nil,
		}
	case "-":
		return Token{
			Position:   lex.cursor,
			LineNumber: lex.lineNumber,
			TokenType:  TOKEN_MINUS,
			Parameter:  nil,
		}
	default:
		tokenInt, _ := strconv.Atoi(text)
		return Token{
			Position:   lex.cursor,
			LineNumber: lex.lineNumber,
			TokenType:  TOKEN_PUSH,
			Parameter:  tokenInt,
		}
	}
}

func (lex *Lexer) LexLine(text []rune) {
	rune_buffer := []rune{}

	for index, char := range text {
		lex.cursor = index
		if unicode.IsSpace(char) {
			lex.cursor = index
			if len(rune_buffer) > 0 {
				lex.tokens = append(lex.tokens, lex.TextToToken(string(rune_buffer)))
				rune_buffer = []rune{}
				continue
			} else {
				continue
			}
		} else {
			rune_buffer = append(rune_buffer, char)
			if index == len(text)-1 {
				lex.tokens = append(lex.tokens, lex.TextToToken(string(rune_buffer)))
				rune_buffer = []rune{}
			}
			continue
		}
	}
}

func (lex *Lexer) Lex() []Token {
	for index, line := range lex.program {
		lex.lineNumber = index
		lex.LexLine(line)
	}

	return lex.tokens
}
