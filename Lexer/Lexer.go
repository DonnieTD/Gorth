package lexer

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"unicode"

	utils "github.com/DonnieTD/NAH/Utils"
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

func (lex *Lexer) CrossReferenceProgram() {
	if COUNT_TOKENS != 7 {
		abs, err := filepath.Abs("./Lexer/Lexer.go")
		if err == nil {
			fmt.Printf("Error in: %v\nUpdate CURRENT_OPCOUNT CrossReferenceProgran() NOTE ONLY BLOCKS NEED TO BE REFERENCED HERE IF ITS NOT A BLOCK INCREMENT AND MOVE ON \n", abs)
		}
		os.Exit(1)
	}

	var block_reference_stack utils.Stack

	for index, token := range lex.Tokens {
		if token.TokenType == TOKEN_IF {
			block_reference_stack.Push(index)
		} else if token.TokenType == TOKEN_END {
			if_addr, _ := block_reference_stack.Pop()
			if_token_index := int((if_addr).(int))
			lex.Tokens[if_token_index] = Token{
				Position:   lex.Tokens[if_token_index].Position,
				LineNumber: lex.Tokens[if_token_index].LineNumber,
				TokenType:  TOKEN_IF,
				// set the if parameter to the address of the end block
				Parameter: index,
			}
		}
	}
}

func (lex *Lexer) TextToToken(text string) Token {
	if COUNT_TOKENS != 7 {
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
	case "if":
		return Token{
			Position:   lex.Cursor,
			LineNumber: lex.LineNumber,
			TokenType:  TOKEN_IF,
			Parameter:  nil,
		}
	case "end":
		return Token{
			Position:   lex.Cursor,
			LineNumber: lex.LineNumber,
			TokenType:  TOKEN_END,
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
	lex.CrossReferenceProgram()
}
