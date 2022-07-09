package nahi

import (
	"fmt"
	"reflect"

	lexer "github.com/DonnieTD/NAH/Lexer"
	utils "github.com/DonnieTD/NAH/Utils"
)

func (n *NAH) Interpret() {
	utils.CountTokensCheck(lexer.COUNT_TOKENS, 15, "./NAHI/NAHI.go:177", "Interpret")

	var programStack utils.Stack
	var programMemory [MEM_CAPACITY]byte

	for t_token_index := 0; t_token_index < len(n.LEXER.Tokens); {
		token := n.LEXER.Tokens[t_token_index]
		switch token.TokenType {
		case lexer.TOKEN_PUSH:
			programStack.Push(token.Parameter)
		case lexer.TOKEN_PLUS:
			a, _ := programStack.Pop()
			b, _ := programStack.Pop()
			if reflect.TypeOf(a).Kind() == reflect.Int && reflect.TypeOf(b).Kind() == reflect.Int {
				a := a.(int)
				b := b.(int)
				programStack.Push(a + b)
			} else if reflect.TypeOf(a).Kind() == reflect.Uint8 && reflect.TypeOf(b).Kind() == reflect.Uint8 {
				a := a.(uint8)
				b := b.(uint8)
				programStack.Push(a + b)
			} else if reflect.TypeOf(a).Kind() == reflect.Int && reflect.TypeOf(b).Kind() == reflect.Uint8 {
				a := uint8(a.(int))
				b := b.(uint8)
				programStack.Push(a + b)
			} else if reflect.TypeOf(a).Kind() == reflect.Uint8 && reflect.TypeOf(b).Kind() == reflect.Int {
				a := a.(uint8)
				b := uint8(b.(int))
				programStack.Push(a + b)
			}
			// later on do string concat here maybe
		case lexer.TOKEN_MINUS:
			a, _ := programStack.Pop()
			b, _ := programStack.Pop()
			if reflect.TypeOf(a).Kind() == reflect.Int && reflect.TypeOf(b).Kind() == reflect.Int {
				a := a.(int)
				b := b.(int)
				programStack.Push(b - a)
			}
		case lexer.TOKEN_EQUALS:
			a, _ := programStack.Pop()
			b, _ := programStack.Pop()
			if reflect.TypeOf(a).Kind() == reflect.Int && reflect.TypeOf(b).Kind() == reflect.Int {
				a := a.(int)
				b := b.(int)
				if a == b {
					programStack.Push(1)
				} else {
					programStack.Push(0)
				}
			}
		case lexer.TOKEN_IF:
			a, _ := programStack.Pop()
			// if false jump to end
			if a == 0 {
				t_token_index = token.Parameter.(int)
				continue
			}
			// otherwise continue executing?
		case lexer.TOKEN_END:
			t_token_index = token.Parameter.(int)
			continue
		case lexer.TOKEN_ELSE:
			t_token_index = token.Parameter.(int)
			continue
		case lexer.TOKEN_DUMP:
			a, _ := programStack.Pop()
			fmt.Printf("%v \n", a)
		case lexer.TOKEN_DUP:
			a, _ := programStack.Pop()
			programStack.Push(a)
			programStack.Push(a)
		case lexer.TOKEN_GREATER_THAN:
			b, _ := programStack.Pop()
			a, _ := programStack.Pop()
			if a.(int) > b.(int) {
				programStack.Push(1)
			} else {
				programStack.Push(0)
			}
		case lexer.TOKEN_WHILE:
			t_token_index++
			continue
		case lexer.TOKEN_DO:
			a, _ := programStack.Pop()

			if a.(int) == 0 {
				t_token_index = token.Parameter.(int)
				continue
			} else {
				t_token_index++
				continue
			}
		case lexer.TOKEN_MEM:
			// MEMORY INDEX STARTS AT ZERO HERE
			programStack.Push(0)
			t_token_index++
			continue
		case lexer.TOKEN_STORE:
			bytee, _ := programStack.Pop()
			addr, _ := programStack.Pop()
			if reflect.TypeOf(addr).Kind() == reflect.Int {
				if reflect.TypeOf(bytee).Kind() == reflect.Int {
					programMemory[addr.(int)] = uint8(bytee.(int)) % 0xFF
				} else {
					programMemory[addr.(int)] = bytee.(uint8) % 0xFF
				}
			} else {
				if reflect.TypeOf(bytee).Kind() == reflect.Int {

					programMemory[int(addr.(uint8))] = uint8(bytee.(int)) % 0xFF
				} else {
					programMemory[int(addr.(uint8))] = bytee.(uint8) % 0xFF
				}
			}
			t_token_index++
			continue
		case lexer.TOKEN_LOAD:
			addr, _ := programStack.Pop()

			bytee := programMemory[addr.(int)]
			programStack.Push(bytee)
			t_token_index++
			continue
		default:
			fmt.Println("Unreachable")
		}
		t_token_index++
	}
	// fmt.Println(programMemory[0:100])
	// fmt.Println(string(programMemory[0:100]))
}
