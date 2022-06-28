package opcreators

import (
	optypes "github.com/DonnieTD/Gorth/OpTypes"
	utils "github.com/DonnieTD/Gorth/Utils"
)

func Push(x interface{}, lineNumber string, tokenPosition string) utils.Tuple {
	return utils.Tuple{
		Optype:        optypes.OP_PUSH,
		Parameters:    x,
		LineNumber:    lineNumber,
		TokenPosition: tokenPosition,
	}
}

func Plus(lineNumber string, tokenPosition string) utils.Tuple {
	return utils.Tuple{
		Optype:        optypes.OP_PLUS,
		Parameters:    nil,
		LineNumber:    lineNumber,
		TokenPosition: tokenPosition,
	}
}

func Minus(lineNumber string, tokenPosition string) utils.Tuple {
	return utils.Tuple{
		Optype:        optypes.OP_MINUS,
		Parameters:    nil,
		LineNumber:    lineNumber,
		TokenPosition: tokenPosition,
	}
}

func Dump(lineNumber string, tokenPosition string) utils.Tuple {
	return utils.Tuple{
		Optype:        optypes.OP_DUMP,
		Parameters:    nil,
		LineNumber:    lineNumber,
		TokenPosition: tokenPosition,
	}
}
