package lexer

import "strings"

func (lex *Lexer) Lex() {
	for index, line := range lex.Program {
		lex.LineNumber = index
		// add single line comments
		if strings.Contains(string(line), "//") {
			lineWithoutComment := []rune(strings.Split(string(line), "//")[0])
			lex.LexLine(lineWithoutComment)
		} else {
			lex.LexLine(line)
		}
	}
	lex.CrossReferenceProgram()
}
