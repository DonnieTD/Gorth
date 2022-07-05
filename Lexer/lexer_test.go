package lexer

import (
	"testing"
)

func TestLexer(t *testing.T) {
	test_lexer := New("../TestScripts/collection.nah")

	if test_lexer.Cursor != 0 {
		t.Errorf("got %q, wanted %q", test_lexer.Cursor, 0)
	}

	if test_lexer.LineNumber != 0 {
		t.Errorf("got %q, wanted %q", test_lexer.LineNumber, 0)
	}

	if test_lexer.FilePath != "../TestScripts/collection.nah" {
		t.Errorf("got %v, wanted %v", test_lexer.FilePath, "../TestScripts/collection.nah")
	}

	if len(test_lexer.Program) != 19 {
		t.Errorf("got %v, wanted %v", len(test_lexer.Program), 19)
	}

	if len(test_lexer.Tokens) != 0 {
		t.Errorf("got %v, wanted %v", len(test_lexer.Tokens), 0)
	}

	test_lexer.LoadProgram()
	test_lexer.Lex()

	if test_lexer.Cursor != 2 {
		t.Errorf("got %q, wanted %q", test_lexer.Cursor, 2)
	}

	if test_lexer.LineNumber != 18 {
		t.Errorf("got %q, wanted %q", test_lexer.LineNumber, 18)
	}

	if len(test_lexer.Program) != 19 {
		t.Errorf("got %v, wanted %v", len(test_lexer.Program), 19)
	}

	if len(test_lexer.Tokens) != 51 {
		t.Errorf("got %v, wanted %v", len(test_lexer.Tokens), 51)
	}
}
