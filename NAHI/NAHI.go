package nahi

import (
	lexer "github.com/DonnieTD/NAH/Lexer"
)

type NAH struct {
	LEXER *lexer.Lexer
}

const MEM_CAPACITY = 640_000
