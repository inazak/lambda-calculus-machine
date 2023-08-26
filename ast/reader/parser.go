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

func (p *Parser) GetError() []string {
	return p.errors
}

func (p *Parser) readToken() {
	p.currToken = p.nextToken
	p.nextToken = p.l.nextToken()

	if p.currToken.Type == UNKNOWN {
		p.errors = append(p.errors, fmt.Sprintf("Lexer got unknown token:"))
		p.errors = append(p.errors, fmt.Sprintf("  %s", p.l.GetError()))
	}
}

func (p *Parser) addErrorMessage(process string, expect string) {
	p.errors = append(p.errors, fmt.Sprintf("Error at Parsing %s:", process))
	p.errors = append(p.errors, fmt.Sprintf("  reading text '%s' <-", shorten(p.l.lookbackText(), 10)))
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

func (p *Parser) isTokenType(expect TokenType) bool {
	return p.currToken.Type == expect
}

func (p *Parser) popToken() Token {
	tk := p.currToken
	p.readToken()
	return tk
}

func (p *Parser) consumeToken() {
	_ = p.popToken()
}

func (p *Parser) Parse() ast.Expression {
	expr := p.parseExpression()

	if !p.isTokenType(EOT) {
		p.addErrorMessage("Parse", "unexpected EOT")
		return expr
	}

	p.consumeToken()
	return expr
}

func (p *Parser) parseExpression() ast.Expression {
	var expr ast.Expression

	switch p.currToken.Type {
	case LAMBDA:
		expr = p.parseFunction()
	case LPAREN:
		expr = p.parseApplication()
	case SYMBOL:
		expr = p.parseSymbol()
	default:
		p.addErrorMessage("Expression", "one of the allowed char")
	}
	return expr
}

func (p *Parser) parseFunction() ast.Function {
	f := ast.Function{}

	if !p.isTokenType(LAMBDA) {
		p.addErrorMessage("Function", "lambda '^' char")
		return f
	}
	p.consumeToken()

	tk := p.popToken()
	if len(tk.Literal) != 1 {
		p.addErrorMessage("Function", "symbol character")
		return f
	}

	f.Arg = tk.Literal

	if !p.isTokenType(DOT) {
		p.addErrorMessage("Function", "dot '.' char")
		return f
	}
	p.consumeToken()

	f.Body = p.parseExpression()
	return f
}

func (p *Parser) parseApplication() ast.Application {
	a := ast.Application{}

	if !p.isTokenType(LPAREN) {
		p.addErrorMessage("Application", "left paren '(' char")
		return a
	}
	p.consumeToken()

	a.Left = p.parseExpression()
	a.Right = p.parseExpression()

	if !p.isTokenType(RPAREN) {
		p.addErrorMessage("Application", "right paren ')' char")
		return a
	}
	p.consumeToken()

	return a
}

func (p *Parser) parseSymbol() ast.Symbol {
	s := ast.Symbol{}

	if !p.isTokenType(SYMBOL) {
		p.addErrorMessage("Symbol", "symbol character")
		return s
	}

	tk := p.popToken()
	if len(tk.Literal) != 1 {
		p.addErrorMessage("Symbol", "symbol character")
		return s
	}
	s.Name = tk.Literal

	return s
}
