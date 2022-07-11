package lexer

import (
	"fmt"
	"os"
	"strconv"

	utils "github.com/DonnieTD/NAH/Utils"
)

func (lex *Lexer) TextToToken(text string) Token {
	utils.CountTokensCheck(COUNT_TOKENS, 17, "./Lexer/TextToToken.go", "TextToToken")
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
	case "syscall3":
		return Token{
			Position:   lex.Cursor,
			LineNumber: lex.LineNumber,
			TokenType:  TOKEN_SYSCALL3,
			Parameter:  nil,
		}
	case "syscall1":
		return Token{
			Position:   lex.Cursor,
			LineNumber: lex.LineNumber,
			TokenType:  TOKEN_SYSCALL1,
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
