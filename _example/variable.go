package main

import (
	"github.com/inazak/lambda-calculus-machine/ast"
)

type Variable struct {
	Name string
	Body ast.Expression
}

func (v Variable) ExpressionString() string {
	return v.Name
}

func expandAll(expr ast.Expression) (ast.Expression) {
	for hasVariable := true ; hasVariable ; hasVariable = false {
	  expr = expand(expr, &hasVariable)
	}
	return expr
}

func expand(expr ast.Expression, flag *bool) ast.Expression {
	switch e := expr.(type) {
	case ast.Symbol:
		return e
	case ast.Application:
		return ast.Application{ Left: expand(e.Left, flag), Right: expand(e.Right, flag) }
	case ast.Function:
		return ast.Function{ Arg: e.Arg, Body: expand(e.Body, flag) }
	case Variable:
		*flag = true
		return e.Body
	default:
		panic("unknown Expression")
	}
}

var VAR_TRUE =
	ast.Function{ Arg: "x", Body:
	ast.Function{ Arg: "y", Body:
	ast.Symbol{ Name: "x" } } }

var VAR_FALSE =
  ast.Function { Arg: "x", Body:
  ast.Function { Arg: "y", Body:
  ast.Symbol { Name: "y" } } }

var VAR_IF =
  ast.Function { Arg: "b", Body:
  ast.Symbol { Name: "b" } }

var VAR_ZERO =
  ast.Function { Arg: "f", Body:
  ast.Function { Arg: "x", Body:
  ast.Symbol { Name: "x" } } }

var VAR_ISZERO =
  ast.Function { Arg: "x", Body:
  ast.Application { Left:
  ast.Application { Left:
  ast.Symbol { Name: "x" }, Right:
  ast.Application { Left: VAR_TRUE, Right: VAR_FALSE, } }, Right: VAR_TRUE, } }

var VAR_LESSOREQ =
  ast.Function { Arg: "x", Body:
  ast.Function { Arg: "y", Body:
  ast.Application { Left: VAR_ISZERO, Right:
  ast.Application { Left:
  ast.Application { Left: VAR_SUB, Right:
  ast.Symbol { Name: "x" } }, Right:
  ast.Symbol { Name: "y" } } } } }

var VAR_SUCC =
  ast.Function { Arg: "n", Body:
  ast.Function { Arg: "f", Body:
  ast.Function { Arg: "x", Body:
  ast.Application { Left:
  ast.Symbol { Name: "f" }, Right:
  ast.Application { Left:
  ast.Application { Left:
  ast.Symbol { Name: "n" }, Right:
  ast.Symbol { Name: "f" } }, Right:
  ast.Symbol { Name: "x" } } } } } }

var VAR_PRED =
  ast.Function { Arg: "n", Body:
  ast.Function { Arg: "f", Body:
  ast.Function { Arg: "x", Body:
  ast.Application { Left:
  ast.Application { Left:
  ast.Application { Left:
  ast.Symbol { Name: "n" }, Right:
  ast.Function { Arg: "g", Body:
  ast.Function { Arg: "h", Body:
  ast.Application { Left:
  ast.Symbol { Name: "h" }, Right:
  ast.Application { Left:
  ast.Symbol { Name: "g" }, Right:
  ast.Symbol { Name: "f" } } } } } }, Right:
  ast.Function { Arg: "u", Body:
  ast.Symbol { Name: "x" } } }, Right:
  ast.Function { Arg: "u", Body:
  ast.Symbol { Name: "u" } } } } } }

var VAR_ADD =
  ast.Function { Arg: "x", Body:
  ast.Function { Arg: "y", Body:
  ast.Application { Left:
  ast.Application { Left:
  ast.Symbol { Name: "x" }, Right: VAR_SUCC, }, Right:
  ast.Symbol { Name: "y" } } } }

var VAR_SUB =
  ast.Function { Arg: "x", Body:
  ast.Function { Arg: "y", Body:
  ast.Application { Left:
  ast.Application { Left:
  ast.Symbol { Name: "y" }, Right: VAR_PRED, }, Right:
  ast.Symbol { Name: "x" } } } }

var VAR_Y =
  ast.Function { Arg: "f", Body:
  ast.Application { Left:
  ast.Function { Arg: "x", Body:
  ast.Application { Left:
  ast.Symbol { Name: "f" }, Right:
  ast.Application { Left:
  ast.Symbol { Name: "x" }, Right:
  ast.Symbol { Name: "x" } } } }, Right:
  ast.Function { Arg: "x", Body:
  ast.Application { Left:
  ast.Symbol { Name: "f" }, Right:
  ast.Application { Left:
  ast.Symbol { Name: "x" }, Right:
  ast.Symbol { Name: "x" } } } } } }

var VAR_DIV =
  ast.Application { Left: VAR_Y, Right:
  ast.Function { Arg: "f", Body:
  ast.Function { Arg: "m", Body:
  ast.Function { Arg: "n", Body:
  ast.Application { Left:
  ast.Application { Left:
  ast.Application { Left: VAR_IF, Right:
  ast.Application { Left:
  ast.Application { Left: VAR_LESSOREQ, Right:
  ast.Symbol { Name: "n" } }, Right:
  ast.Symbol { Name: "m" } } }, Right:
  ast.Application { Left: VAR_SUCC, Right:
  ast.Application { Left:
  ast.Application { Left:
  ast.Symbol { Name: "f" }, Right:
  ast.Application { Left:
  ast.Application { Left: VAR_SUB, Right:
  ast.Symbol { Name: "m" } }, Right:
  ast.Symbol { Name: "n" } } }, Right:
  ast.Symbol { Name: "n" } } } }, Right: VAR_ZERO, } } } } }

var VAR_MOD =
  ast.Application { Left: VAR_Y, Right:
  ast.Function { Arg: "f", Body:
  ast.Function { Arg: "m", Body:
  ast.Function { Arg: "n", Body:
  ast.Application { Left:
  ast.Application { Left:
  ast.Application { Left: VAR_IF, Right:
  ast.Application { Left:
  ast.Application { Left: VAR_LESSOREQ, Right:
  ast.Symbol { Name: "n" } }, Right:
  ast.Symbol { Name: "m" } } }, Right:
  ast.Application { Left:
  ast.Application { Left:
  ast.Symbol { Name: "f" }, Right:
  ast.Application { Left:
  ast.Application { Left: VAR_SUB, Right:
  ast.Symbol { Name: "m" } }, Right:
  ast.Symbol { Name: "n" } } }, Right:
  ast.Symbol { Name: "n" } } }, Right:
  ast.Symbol { Name: "m" } } } } } }


