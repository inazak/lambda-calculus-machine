package main

import (
	"github.com/inazak/lambda-calculus-machine/ast"
)

// church encoding
// 0 = ^f.^x.x
// 1 = ^f.^x.(f x)
// 2 = ^f.^x.(f (f x))
func Number(n int) ast.Function {
	var expr ast.Expression

	expr = ast.Symbol{Name: "x"}
	for ; n > 0; n -= 1 {
		expr = ast.Application{
			Left:  ast.Symbol{Name: "f"},
			Right: expr,
		}
	}

	expr = ast.Function{
		Arg: "f",
		Body: ast.Function{
			Arg:  "x",
			Body: expr,
		},
	}

	f, _ := expr.(ast.Function)
	return f
}
