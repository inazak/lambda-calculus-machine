package reader

type Lexer struct {
	text         []rune
	line         int
	currPosition int
	nextPosition int
	c            rune //charactor of currPosition
	msg          string
}

func NewLexer(text string) *Lexer {
	l := &Lexer{text: []rune(text), line: 1}
	l.read()
	return l
}

func (l *Lexer) LookbackText() string {
	if l.currPosition < 2 {
		return ""
	} else {
		return string(l.text[:l.currPosition-2])
	}
}

func (l *Lexer) GetMsg() string {
	return l.msg
}

func (l *Lexer) read() {
	if l.nextPosition >= len(l.text) {
		l.c = 0
	} else {
		l.c = l.text[l.nextPosition]
	}
	l.currPosition = l.nextPosition
	l.nextPosition += 1
}

func (l *Lexer) NextToken() Token {
	var tk Token
	l.skipSpace()

	switch l.c {
	case '^':
		tk = NewToken(LAMBDA, "^")
	case '.':
		tk = NewToken(DOT, ".")
	case '(':
		tk = NewToken(LPAREN, "(")
	case ')':
		tk = NewToken(RPAREN, ")")
	case 0:
		tk = NewToken(EOT, "")
	default:
		if l.isSymbol() {
			tk = NewToken(SYMBOL, string(l.c))
		} else {
			tk = NewToken(UNKNOWN, string(l.c))
			l.msg = "unallowed character -> " + string(l.c)
		}
	}

	l.read()
	return tk
}

func (l *Lexer) skipSpace() {
	for l.c == ' ' || l.c == '\t' || l.c == '\r' || l.c == '\n' {
		l.read()
		if l.c == '\n' {
			l.line += 1
		}
	}
}

func (l *Lexer) nextIs(r rune) bool {
	if l.nextPosition < len(l.text) {
		c := l.text[l.nextPosition]
		if c == r {
			return true
		}
	}
	return false
}

func (l *Lexer) isSymbol() bool {
	return 'a' <= l.c && l.c <= 'z'
}
