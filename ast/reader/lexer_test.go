package reader

import (
	"testing"
)

func TestNextToken1(t *testing.T) {

	text := "^x.^y.((a b) c)"
	l := NewLexer(text)

	expected := []struct {
		Type    TokenType
		Literal string
	}{
		{Type: LAMBDA, Literal: "^"},
		{Type: SYMBOL, Literal: "x"},
		{Type: DOT, Literal: "."},
		{Type: LAMBDA, Literal: "^"},
		{Type: SYMBOL, Literal: "y"},
		{Type: DOT, Literal: "."},
		{Type: LPAREN, Literal: "("},
		{Type: LPAREN, Literal: "("},
		{Type: SYMBOL, Literal: "a"},
		{Type: SYMBOL, Literal: "b"},
		{Type: RPAREN, Literal: ")"},
		{Type: SYMBOL, Literal: "c"},
		{Type: RPAREN, Literal: ")"},
		{Type: EOT, Literal: ""},
	}

	for i, e := range expected {
		tk := l.nextToken()
		if tk.Type != e.Type {
			t.Fatalf("No.%v expected type=%q, got=%q", i, e.Type, tk.Type)
		}
		if tk.Literal != e.Literal {
			t.Fatalf("No.%v expected literal=%q, got=%q", i, e.Literal, tk.Literal)
		}
	}
}
