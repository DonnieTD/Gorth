package main

import (
	"fmt"
	"reflect"

	utils "github.com/DonnieTD/Gorth/Utils"
)

func SimulateProgram(program []utils.Tuple) {
	if COUNT_OPS != 4 {
		errors.New("Update CURRENT_OPCOUNT SimulateProgram")
		return
	}

	var programstack utils.Stack

	for _, operation := range program {
		switch operation.Optype {
		case OP_PUSH:
			programstack.Push(operation.Parameters)
		case OP_PLUS:
			a, _ := programstack.Pop()
			b, _ := programstack.Pop()
			if reflect.TypeOf(a).Kind() == reflect.Int && reflect.TypeOf(b).Kind() == reflect.Int {
				a := a.(int)
				b := b.(int)
				programstack.Push(a + b)
			}
			// later on do string concat here maybe
		case OP_MINUS:
			a, _ := programstack.Pop()
			b, _ := programstack.Pop()
			if reflect.TypeOf(a).Kind() == reflect.Int && reflect.TypeOf(b).Kind() == reflect.Int {
				a := a.(int)
				b := b.(int)
				programstack.Push(b - a)
			}
		case OP_DUMP:
			a, _ := programstack.Pop()
			fmt.Printf("%v \n", a)
		default:
			fmt.Println("Unreachable")
		}
	}
}
