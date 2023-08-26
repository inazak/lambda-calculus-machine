package reader

type TokenType string

const (
	SYMBOL  = "SYMBOL"
	LAMBDA  = "LAMBDA"
	DOT     = "DOT"
	LPAREN  = "LPAREN"
	RPAREN  = "RPAREN"
	EOT     = "EOT"
	UNKNOWN = "UNKNOWN"
)

type Token struct {
	Type    TokenType
	Literal string
}

func NewToken(t TokenType, s string) Token {
	return Token{Type: t, Literal: s}
}
