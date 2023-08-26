package vm

import (
	"github.com/inazak/lambda-calculus-machine/ast"
	"testing"
)

func TestCompile1(t *testing.T) {

	expr := ast.Symbol{Name: "x"}
	expect := []Instruction{Fetch{Name: "x"}}

	code := Compile(expr)

	for i, inst := range code {
		got := inst.InstructionString()
		want := expect[i].InstructionString()
		if got != want {
			t.Errorf("expect=%s, but got=%s", want, got)
		}
	}
}
