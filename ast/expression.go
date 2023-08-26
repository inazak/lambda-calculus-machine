package ast

type Expression interface {
	ExpressionString() string
}

type Symbol struct {
	Name string
}

type Function struct {
	Arg  string
	Body Expression
}

type Application struct {
	Left  Expression
	Right Expression
}

func (s Symbol) ExpressionString() string {
	return s.Name
}

func (f Function) ExpressionString() string {
	return "^" + f.Arg + "." + f.Body.ExpressionString()
}

func (a Application) ExpressionString() string {
	return "(" + a.Left.ExpressionString() + " " + a.Right.ExpressionString() + ")"
}
