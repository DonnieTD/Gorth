package main

import (
	"fmt"
	"os"
	"reflect"

	lexer "github.com/DonnieTD/Gorth/Lexer"
	utils "github.com/DonnieTD/Gorth/Utils"
)

func SimulateProgram(program []lexer.Token) {
	fmt.Printf("%v \n", program)
	if lexer.COUNT_TOKENS != 4 {
		fmt.Println("Update CURRENT_OPCOUNT SimulateProgram")
		os.Exit(1)
	}

	var programstack utils.Stack

	for _, token := range program {
		switch token.TokenType {
		case lexer.TOKEN_PUSH:
			programstack.Push(token.Parameter)
		case lexer.TOKEN_PLUS:
			a, _ := programstack.Pop()
			b, _ := programstack.Pop()
			if reflect.TypeOf(a).Kind() == reflect.Int && reflect.TypeOf(b).Kind() == reflect.Int {
				a := a.(int)
				b := b.(int)
				programstack.Push(a + b)
			}
			// later on do string concat here maybe
		case lexer.TOKEN_MINUS:
			a, _ := programstack.Pop()
			b, _ := programstack.Pop()
			if reflect.TypeOf(a).Kind() == reflect.Int && reflect.TypeOf(b).Kind() == reflect.Int {
				a := a.(int)
				b := b.(int)
				programstack.Push(b - a)
			}
		case lexer.TOKEN_DUMP:
			a, _ := programstack.Pop()
			fmt.Printf("%v \n", a)
		default:
			fmt.Println("Unreachable")
		}
	}
}
