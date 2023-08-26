package reader

import (
	"testing"
)

func TestParse1(t *testing.T) {

	text := "^a.^b.((^c.c d) e)"

	l := NewLexer(text)
	p := NewParser(l)
	expr := p.Parse()

	if errmsg := p.GetError(); errmsg != nil {
		t.Fatalf("%s", errmsg)
	}

	if text != expr.ExpressionString() {
		t.Errorf("expected literal=%q, got=%q", text, expr.ExpressionString())
	}
}
