package lexer

import "unicode"

func (lex *Lexer) LexLine(text []rune) {
	rune_buffer := []rune{}

	for index, char := range text {
		lex.Cursor = index
		if unicode.IsSpace(char) {
			lex.Cursor = index
			if len(rune_buffer) > 0 {
				lex.Tokens = append(lex.Tokens, lex.TextToToken(string(rune_buffer)))
				rune_buffer = []rune{}
				continue
			} else {
				continue
			}
		} else {
			rune_buffer = append(rune_buffer, char)
			if index == len(text)-1 {
				lex.Tokens = append(lex.Tokens, lex.TextToToken(string(rune_buffer)))
				rune_buffer = []rune{}
			}
			continue
		}
	}
}
