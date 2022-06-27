package opcreators

import (
	optypes "github.com/DonnieTD/Gorth/OpTypes"
	utils "github.com/DonnieTD/Gorth/Utils"
)

func Push(x interface{}) utils.Tuple {
	return utils.Tuple{
		Optype:     optypes.OP_PUSH,
		Parameters: x,
	}
}

func Plus() utils.Tuple {
	return utils.Tuple{
		Optype:     optypes.OP_PLUS,
		Parameters: nil,
	}
}

func Minus() utils.Tuple {
	return utils.Tuple{
		Optype:     optypes.OP_MINUS,
		Parameters: nil,
	}
}

func Dump() utils.Tuple {
	return utils.Tuple{
		Optype:     optypes.OP_DUMP,
		Parameters: nil,
	}
}
