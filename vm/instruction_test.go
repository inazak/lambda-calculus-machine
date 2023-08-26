package vm

import (
	"github.com/inazak/lambda-calculus-machine/ast"
	"github.com/inazak/lambda-calculus-machine/ast/reader"
	"testing"
)

func parseText(text string) ast.Expression {
	l := reader.NewLexer(text)
	p := reader.NewParser(l)
	return p.Parse()
}

func TestCompile1(t *testing.T) {

	text := "x"
	expect := []Instruction{
		Fetch{Name: "x"},
	}

	code := Compile(parseText(text))

	for i, inst := range code {
		got := inst.InstructionString()
		exp := expect[i].InstructionString()
		if got != exp {
			t.Errorf("expect=%s, but got=%s", exp, got)
		}
	}
}

func TestCompile2(t *testing.T) {

	text := "(^x.x y)"
	expect := []Instruction{
		Call{
			Code: []Instruction{
				Close{Arg: "x", Code: []Instruction{Fetch{Name: "x"}, Return{}}},
				Fetch{Name: "y"},
				Apply{},
				Return{},
			},
		},
	}

	code := Compile(parseText(text))

	for i, inst := range code {
		got := inst.InstructionString()
		exp := expect[i].InstructionString()
		if got != exp {
			t.Errorf("expect=%s, but got=%s", exp, got)
		}
	}
}
