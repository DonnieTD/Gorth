package main

import utils "github.com/DonnieTD/Gorth/Utils"

func Push(x interface{}) utils.Tuple {
	return utils.Tuple{
		Optype:     OP_PUSH,
		Parameters: x,
	}
}

func Plus() utils.Tuple {
	return utils.Tuple{
		Optype:     OP_PLUS,
		Parameters: nil,
	}
}

func Minus() utils.Tuple {
	return utils.Tuple{
		Optype:     OP_MINUS,
		Parameters: nil,
	}
}

func Dump() utils.Tuple {
	return utils.Tuple{
		Optype:     OP_DUMP,
		Parameters: nil,
	}
}
