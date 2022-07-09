package lexer

func New(filePath string) *Lexer {
	lex := Lexer{
		// LINE NUMBERS AND COLS ARE 1 INDEXED IN REPORTING REMEMBER TO INCREMENT
		Cursor:     0,
		LineNumber: 0,
		FilePath:   filePath,
		Program:    [][]rune{},
		Tokens:     []Token{},
	}

	lex.LoadProgram()

	return &lex
}
