package lexer

type Lexer struct {
	Cursor     int
	LineNumber int
	FilePath   string
	Program    [][]rune
	Tokens     []Token
}
