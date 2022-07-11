package lexer

import utils "github.com/DonnieTD/NAH/Utils"

// takes in a program
// when it finds an if it pushes its position to block_reference_stack
//	when it finds else it takes the last if address off of the stack
//	it then finds that specific if and adds the
func (lex *Lexer) CrossReferenceProgram() {
	// NOTE ONLY BLOCKS NEED TO BE REFERENCED HERE IF ITS NOT A BLOCK INCREMENT AND MOVE ON
	utils.CountTokensCheck(COUNT_TOKENS, 17, "./Lexer/CrossReferenceProgram.go", "CrossReferenceProgram")

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
