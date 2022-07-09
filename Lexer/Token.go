package lexer

type Token struct {
	TokenType  int
	Parameter  interface{}
	LineNumber int
	Position   int
}
