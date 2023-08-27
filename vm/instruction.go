package vm

import (
	"github.com/inazak/lambda-calculus-machine/ast"
	"strings"
)

type Instruction interface {
	InstructionString() string
}

type Fetch struct {
	Name string
}

type Call struct {
	Code []Instruction
}

type Close struct {
	Arg  string
	Code []Instruction
}

type Apply struct {
	//nothing
}

type Return struct {
	//nothing
}

type Abstract struct {
	//nothing
}

type Swap struct {
	//nothing
}

func (f Fetch) InstructionString() string {
	return "Fetch " + f.Name
}

func (c Call) InstructionString() string {
	list := []string{}
	for _, code := range c.Code {
		list = append(list, code.InstructionString())
	}
	return "Call [" + strings.Join(list, "; ") + "]"
}

func (c Close) InstructionString() string {
	list := []string{}
	for _, code := range c.Code {
		list = append(list, code.InstructionString())
	}
	return "Close " + c.Arg + ", [" + strings.Join(list, "; ") + "]"
}

func (a Apply) InstructionString() string {
	return "Apply"
}

func (r Return) InstructionString() string {
	return "Return"
}

func (a Abstract) InstructionString() string {
	return "Abstract"
}

func (w Swap) InstructionString() string {
	return "Swap"
}

func Compile(expr ast.Expression) []Instruction {

	switch e := expr.(type) {

	case ast.Symbol:
		return []Instruction{
			Fetch{Name: e.Name},
		}

	case ast.Application:
		left := Compile(e.Left)
		right := Compile(e.Right)
		return []Instruction{
			Call{
				Code: append(append(left, right...), Apply{}, Return{}),
			},
		}

	case ast.Function:
		return []Instruction{
			Close{
				Arg:  e.Arg,
				Code: append(Compile(e.Body), Return{}),
			},
		}

	default:
		panic("compile: unknown expression")
	}
}
