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
	var tok token.Token

	switch l.ch {
	case ',':
		tok = token.New(token.COMMA, l.ch)
	case ';':
		tok = token.New(token.SEMICOLON, l.ch)
	case '=':
		tok = token.New(token.ASSIGN, l.ch)
	case '+':
		tok = token.New(token.PLUS, l.ch)
	case '(':
		tok = token.New(token.LPAREN, l.ch)
	case ')':
		tok = token.New(token.RPAREN, l.ch)
	case '}':
		tok = token.New(token.RBRACE, l.ch)
	case '{':
		tok = token.New(token.LBRACE, l.ch)
	case '|':
		if l.follows('>') {
			tok.Literal = "|>"
			tok.Type = token.PIPE
			l.readChar()
		} else {
			tok = token.New(token.ILLEGAL, l.ch)
		}
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.GetIdentifierType(tok.Literal)
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

func (l *Lexer) follows(s byte) bool {
	if l.pos >= len(l.input)-1 {
		return false
	} else {
		return l.input[l.pos+1] == s
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}
