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

func TestSymbol(t *testing.T) {
	text   := "x"
	expect := "x"

	r := readCompileAndRun(t, text)

	result := r.ValueString()
	if result != expect {
		t.Errorf("expected=%s, but got=%s", expect, result)
	}
}

func TestApplication(t *testing.T) {
	text   := "(a b)"
	expect := "(a b)"

	r := readCompileAndRun(t, text)

	result := r.ValueString()
	if result != expect {
		t.Errorf("expected=%s, but got=%s", expect, result)
	}
}

func TestClosure(t *testing.T) {
	text   := "(^x.x (a b))"
	expect := "(a b)"

	r := readCompileAndRun(t, text)

	result := r.ValueString()
	if result != expect {
		t.Errorf("expected=%s, but got=%s", expect, result)
	}
}

func TestSucc(t *testing.T) {
  text   := "(^n.^f.^x.(f ((n f) x)) ^f.^x.(f x))" //(Succ 1)
  expect := "^f.^x.(f (f x))" //2

	r := readCompileAndRun(t, text)

	result := r.ValueString()
	if result != expect {
		t.Errorf("expected=%s, but got=%s", expect, result)
	}
}

func TestPred(t *testing.T) {
  text   := "(^n.^f.^x.(((n ^g.^h.(h (g f))) ^u.x) ^u.u) ^f.^x.(f (f x)))" //(Pred 2)
  expect := "^f.^x.(f x)" //1

	r := readCompileAndRun(t, text)

	result := r.ValueString()
	if result != expect {
		t.Errorf("expected=%s, but got=%s", expect, result)
	}
}

func TestSub(t *testing.T) {
  text   := "((^x.^y.((y ^n.^f.^x.(((n ^g.^h.(h (g f))) ^u.x) ^u.u)) x) ^f.^x.(f (f x))) ^f.^x.(f x))" //((Sub 2) 1)
  expect := "^f.^x.(f x)" //1

	r := readCompileAndRun(t, text)

	result := r.ValueString()
	if result != expect {
		t.Errorf("expected=%s, but got=%s", expect, result)
	}
}

func TestIsZeroT(t *testing.T) {
  text   := "(^x.((x (^x.^y.x ^x.^y.y)) ^x.^y.x) ^f.^x.x)" //(IsZero 0)
  expect := "^x.^y.x" //True

	r := readCompileAndRun(t, text)

	result := r.ValueString()
	if result != expect {
		t.Errorf("expected=%s, but got=%s", expect, result)
	}
}

func TestIsZeroF(t *testing.T) {
  text   := "(^x.((x (^x.^y.x ^x.^y.y)) ^x.^y.x) ^f.^x.(f x))" //(IsZero 1)
  expect := "^x.^y.y" //False

	r := readCompileAndRun(t, text)

	result := r.ValueString()
	if result != expect {
		t.Errorf("expected=%s, but got=%s", expect, result)
	}
}

func TestMod(t *testing.T) {
  text   := "(((^f.(^x.(f (x x)) ^x.(f (x x))) ^f.^m.^n.(((^b.b ((^x.^y.(^x.((x (^x.^y.x ^x.^y.y)) ^x.^y.x) ((^x.^y.((y ^n.^f.^x.(((n ^g.^h.(h (g f))) ^u.x) ^u.u)) x) x) y)) n) m)) ((f ((^x.^y.((y ^n.^f.^x.(((n ^g.^h.(h (g f))) ^u.x) ^u.u)) x) m) n)) n)) m)) ^f.^x.(f (f (f (f (f x)))))) ^f.^x.(f (f x)))" //((Mod 5) 2)
  expect := "^f.^x.(f x)" //1

	r := readCompileAndRun(t, text)

	result := r.ValueString()
	if result != expect {
		t.Errorf("expected=%s, but got=%s", expect, result)
	}
}

