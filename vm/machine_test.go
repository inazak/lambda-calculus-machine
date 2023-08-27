package vm

import (
	"testing"
	"github.com/inazak/lambda-calculus-machine/ast/reader"
)

func readCompileAndRun(t *testing.T, text string) Value {
  l := reader.NewLexer(text)
  p := reader.NewParser(l)
  expr := p.Parse()
  code := Compile(expr)

  vm := NewVM(code)
	result := vm.Run()

  return result
}

func Test1(t *testing.T) {
	text   := "x"
	expect := "x"

	r := readCompileAndRun(t, text)

	result := r.ValueString()
	if result != expect {
		t.Errorf("expected=%s, but got=%s", expect, result)
	}
}

func Test2(t *testing.T) {
	text   := "(a b)"
	expect := "(a b)"

	r := readCompileAndRun(t, text)

	result := r.ValueString()
	if result != expect {
		t.Errorf("expected=%s, but got=%s", expect, result)
	}
}

func Test3(t *testing.T) {
	text   := "(^x.x (a b))"
	expect := "(a b)"

	r := readCompileAndRun(t, text)

	result := r.ValueString()
	if result != expect {
		t.Errorf("expected=%s, but got=%s", expect, result)
	}
}
