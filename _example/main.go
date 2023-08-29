package main

import (
	"fmt"
	"github.com/inazak/lambda-calculus-machine/ast"
	"github.com/inazak/lambda-calculus-machine/vm"
)

// this is ((mod 100) 13) = 100 mod 13 = 9
// output: ^f.^x.(f (f (f (f (f (f (f (f (f x)))))))))
func main() {

	var expr ast.Expression
	expr = ast.Application{Left: ast.Application{Left: VAR_MOD, Right: Number(100)}, Right: Number(13)}
	expr = expandAll(expr)

	code := vm.Compile(expr)
	machine := vm.NewVM(code)
	result := machine.Run()

	fmt.Printf("input: %s\n", expr.ExpressionString())
	fmt.Printf("output: %s\n", result.ValueString())
}
