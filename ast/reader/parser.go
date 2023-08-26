package reader

import (
	"fmt"
	"github.com/inazak/lambda-calculus-machine/ast"
)

type Parser struct {
	l         *Lexer
	currToken Token
	nextToken Token
	errors    []string
}

func NewParser(l *Lexer) *Parser {
	p := &Parser{l: l, errors: nil}
	p.readToken() // set p.nextToken
	p.readToken() // set p.nextToken and p.currToken
	return p
}

func (p *Parser) readToken() {
	p.currToken = p.nextToken
	p.nextToken = p.l.NextToken()

	if p.currToken.Type == UNKNOWN {
		p.errors = append(p.errors, fmt.Sprintf("Lexer got unknown token:"))
		p.errors = append(p.errors, fmt.Sprintf("  %s", p.l.GetMsg()))
	}
}

func (p *Parser) AddErrorMessage(process string, expect string) {
	p.errors = append(p.errors, fmt.Sprintf("Error at Parsing %s:", process))
	p.errors = append(p.errors, fmt.Sprintf("  reading text '%s' <-", shorten(p.l.LookbackText(), 10)))
	p.errors = append(p.errors, fmt.Sprintf("  parser expect %s,", expect))
	p.errors = append(p.errors, fmt.Sprintf("  but got %s '%s'", p.currToken.Type, p.currToken.Literal))
}

func shorten(s string, i int) string {
	if len(s) <= i {
		return s
	} else {
		return "..." + s[len(s)-i:len(s)]
	}
}

func (p *Parser) GetError() []string {
	return p.errors
}

func (p *Parser) IsTokenType(expect TokenType) bool {
	return p.currToken.Type == expect
}

func (p *Parser) PopToken() Token {
	tk := p.currToken
	p.readToken()
	return tk
}

func (p *Parser) ConsumeToken() {
	_ = p.PopToken()
}

func (p *Parser) Parse() ast.Expression {
	expr := p.ParseExpression()

	if !p.IsTokenType(EOT) {
		p.AddErrorMessage("Parse", "unexpected EOT")
		return expr
	}

	p.ConsumeToken()
	return expr
}

func (p *Parser) ParseExpression() ast.Expression {
	var expr ast.Expression

	switch p.currToken.Type {
	case LAMBDA:
		expr = p.ParseFunction()
	case LPAREN:
		expr = p.ParseApplication()
	case SYMBOL:
		expr = p.ParseSymbol()
	default:
		p.AddErrorMessage("Expression", "one of the allowed char")
	}
	return expr
}

func (p *Parser) ParseFunction() ast.Function {
	f := ast.Function{}

	if !p.IsTokenType(LAMBDA) {
		p.AddErrorMessage("Function", "lambda '^' char")
		return f
	}
	p.ConsumeToken()

	tk := p.PopToken()
	if len(tk.Literal) != 1 {
		p.AddErrorMessage("Function", "symbol character")
		return f
	}

	f.Arg = tk.Literal

	if !p.IsTokenType(DOT) {
		p.AddErrorMessage("Function", "dot '.' char")
		return f
	}
	p.ConsumeToken()

	f.Body = p.ParseExpression()
	return f
}

func (p *Parser) ParseApplication() ast.Application {
	a := ast.Application{}

	if !p.IsTokenType(LPAREN) {
		p.AddErrorMessage("Application", "left paren '(' char")
		return a
	}
	p.ConsumeToken()

	a.Left = p.ParseExpression()
	a.Right = p.ParseExpression()

	if !p.IsTokenType(RPAREN) {
		p.AddErrorMessage("Application", "right paren ')' char")
		return a
	}
	p.ConsumeToken()

	return a
}

func (p *Parser) ParseSymbol() ast.Symbol {
	s := ast.Symbol{}

	if !p.IsTokenType(SYMBOL) {
		p.AddErrorMessage("Symbol", "symbol character")
		return s
	}

	tk := p.PopToken()
	if len(tk.Literal) != 1 {
		p.AddErrorMessage("Symbol", "symbol character")
		return s
	}
	s.Name = tk.Literal

	return s
}
