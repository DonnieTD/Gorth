package main

import (
	"fmt"
	"reflect"

	optypes "github.com/DonnieTD/Gorth/OpTypes"
	utils "github.com/DonnieTD/Gorth/Utils"
)

func SimulateProgram(program []utils.Tuple) {
	if optypes.COUNT_OPS != 4 {
		panic("Update CURRENT_OPCOUNT SimulateProgram")
	}

	var programstack utils.Stack

	for _, operation := range program {
		switch operation.Optype {
		case optypes.OP_PUSH:
			programstack.Push(operation.Parameters)
		case optypes.OP_PLUS:
			a, _ := programstack.Pop()
			b, _ := programstack.Pop()
			if reflect.TypeOf(a).Kind() == reflect.Int && reflect.TypeOf(b).Kind() == reflect.Int {
				a := a.(int)
				b := b.(int)
				programstack.Push(a + b)
			}
			// later on do string concat here maybe
		case optypes.OP_MINUS:
			a, _ := programstack.Pop()
			b, _ := programstack.Pop()
			if reflect.TypeOf(a).Kind() == reflect.Int && reflect.TypeOf(b).Kind() == reflect.Int {
				a := a.(int)
				b := b.(int)
				programstack.Push(b - a)
			}
		case optypes.OP_DUMP:
			a, _ := programstack.Pop()
			fmt.Printf("%v \n", a)
		default:
			fmt.Println("Unreachable")
		}
	}
}
