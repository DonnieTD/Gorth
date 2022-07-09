package lexer

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

// takes in a program
// when it finds an if it pushes its position to block_reference_stack
//	when it finds else it takes the last if address off of the stack
//	it then finds that specific if and adds the
func (lex *Lexer) CrossReferenceProgram() {
	// NOTE ONLY BLOCKS NEED TO BE REFERENCED HERE IF ITS NOT A BLOCK INCREMENT AND MOVE ON
	utils.CountTokensCheck(COUNT_TOKENS, 15, "./Lexer/Lexer.go:66", "CrossReferenceProgram")

	var block_reference_stack utils.Stack

	for index, token := range lex.Tokens {
		if token.TokenType == TOKEN_IF {
			block_reference_stack.Push(index)
		} else if token.TokenType == TOKEN_ELSE {
			if_addr, _ := block_reference_stack.Pop()
			if_token_index := int((if_addr).(int))
			lex.Tokens[if_token_index].TokenType = TOKEN_IF
			lex.Tokens[if_token_index].Parameter = index + 1
			block_reference_stack.Push(index)
		} else if token.TokenType == TOKEN_END {
			block_addr, _ := block_reference_stack.Pop()
			block_token_index := int((block_addr).(int))
			if lex.Tokens[block_token_index].TokenType == TOKEN_IF || lex.Tokens[block_token_index].TokenType == TOKEN_ELSE {
				lex.Tokens[block_token_index].Parameter = index
				lex.Tokens[index].Parameter = index + 1
			} else if lex.Tokens[block_token_index].TokenType == TOKEN_DO {
				lex.Tokens[index].Parameter = lex.Tokens[block_token_index].Parameter
				lex.Tokens[block_token_index].Parameter = index + 1
			}
		} else if token.TokenType == TOKEN_WHILE {
			block_reference_stack.Push(index)
		} else if token.TokenType == TOKEN_DO {
			while_ip, _ := block_reference_stack.Pop()
			lex.Tokens[index].Parameter = while_ip
			block_reference_stack.Push(index)
		}
	}
}

func (lex *Lexer) TextToToken(text string) Token {
	utils.CountTokensCheck(COUNT_TOKENS, 15, "./Lexer/Lexer.go:101", "TextToToken")
	switch text {
	case "dump":
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
	case "else":
		return Token{
			Position:   lex.Cursor,
			LineNumber: lex.LineNumber,
			TokenType:  TOKEN_ELSE,
			Parameter:  nil,
		}
	case "end":
		return Token{
			Position:   lex.Cursor,
			LineNumber: lex.LineNumber,
			TokenType:  TOKEN_END,
			Parameter:  nil,
		}
	case "dup":
		return Token{
			Position:   lex.Cursor,
			LineNumber: lex.LineNumber,
			TokenType:  TOKEN_DUP,
			Parameter:  nil,
		}
	case ">":
		return Token{
			Position:   lex.Cursor,
			LineNumber: lex.LineNumber,
			TokenType:  TOKEN_GREATER_THAN,
			Parameter:  nil,
		}
	case "while":
		return Token{
			Position:   lex.Cursor,
			LineNumber: lex.LineNumber,
			TokenType:  TOKEN_WHILE,
			Parameter:  nil,
		}
	case "do":
		return Token{
			Position:   lex.Cursor,
			LineNumber: lex.LineNumber,
			TokenType:  TOKEN_DO,
			Parameter:  nil,
		}
	case "mem":
		return Token{
			Position:   lex.Cursor,
			LineNumber: lex.LineNumber,
			TokenType:  TOKEN_MEM,
			Parameter:  nil,
		}
	case ".":
		return Token{
			Position:   lex.Cursor,
			LineNumber: lex.LineNumber,
			TokenType:  TOKEN_STORE,
			Parameter:  nil,
		}
	case ",":
		return Token{
			Position:   lex.Cursor,
			LineNumber: lex.LineNumber,
			TokenType:  TOKEN_LOAD,
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
		// add single line comments
		if strings.Contains(string(line), "//") {
			lineWithoutComment := []rune(strings.Split(string(line), "//")[0])
			lex.LexLine(lineWithoutComment)
		} else {
			lex.LexLine(line)
		}
	}
	lex.CrossReferenceProgram()
}
