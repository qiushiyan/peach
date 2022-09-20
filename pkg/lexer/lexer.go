package lexer

import "github.com/qiushiyan/peach/pkg/token"

type Lexer struct {
	input   string
	pos     int  // current position in input (points to current char)
	nextpos int  // current reading position in input (after current char)
	ch      byte // current char under examination
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()
	return lexer
}

func (l *Lexer) NextToken() token.Token {
	l.skipWhitespace()

	var tok token.Token

	switch l.ch {
	// single character tokens
	case ',':
		tok = token.New(token.COMMA, l.ch)
	case ';':
		tok = token.New(token.SEMICOLON, l.ch)
	case '=':
		if l.followedBy([]byte{'='}) {
			tok = token.New(token.EQ, "==")
			l.readChar()
		} else {
			tok = token.New(token.ASSIGN, l.ch)
		}
	case '+':
		tok = token.New(token.PLUS, l.ch)
	case '-':
		tok = token.New(token.MINUS, l.ch)
	case '/':
		tok = token.New(token.DIV, l.ch)
	case '*':
		tok = token.New(token.MUL, l.ch)
	case '!':
		if l.followedBy([]byte{'='}) {
			tok = token.New(token.NOT_EQ, "!=")
			l.readChar()
		} else {
			tok = token.New(token.BANG, l.ch)
		}
	case '<':
		if l.followedBy([]byte{'='}) {
			tok = token.New(token.LTE, "<=")
			l.readChar()
		} else {
			tok = token.New(token.LT, l.ch)
		}
	case '>':
		if l.followedBy([]byte{'='}) {
			tok = token.New(token.GTE, ">=")
			l.readChar()
		} else {
			tok = token.New(token.GT, l.ch)
		}
	case '(':
		tok = token.New(token.LPAREN, l.ch)
	case ')':
		tok = token.New(token.RPAREN, l.ch)
	case '}':
		tok = token.New(token.RBRACE, l.ch)
	case '{':
		tok = token.New(token.LBRACE, l.ch)
	// multi character tokens
	case '|':
		if l.followedBy([]byte{'>'}) {
			tok = token.New(token.PIPE, "|>")
			l.readChar()
		} else {
			tok = token.New(token.ILLEGAL, l.ch)
		}
	case 0:
		tok = token.New(token.EOF, "")
	default:
		if isLetter(l.ch) {
			literal := l.readIdentifier()
			tok = token.New(token.GetIdentifierType(literal), literal)
			return tok
		} else if isDigit(l.ch) {
			tok = token.New(token.INTEGER, l.readNumber())
			return tok
		} else {
			tok = token.New(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readChar() {

	if l.nextpos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.nextpos]
	}
	l.pos = l.nextpos
	l.nextpos += 1
}

func (l *Lexer) readIdentifier() string {
	pos := l.pos
	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[pos:l.pos]
}

func (l *Lexer) readNumber() string {
	pos := l.pos
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[pos:l.pos]
}

func (l *Lexer) followedBy(chs []byte) bool {
	if l.pos >= len(l.input)-len(chs) {
		return false
	} else {
		for i, ch := range chs {
			if ch != l.input[l.pos+i+1] {
				return false
			}
		}
		return true
	}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' || l.ch == '\n' { // line break is skipped for now
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
