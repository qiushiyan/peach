package lexer

import (
	"io"
	"text/scanner"

	"github.com/qiushiyan/qlang/pkg/token"
)

type Lexer struct {
	s scanner.Scanner

	ch rune
}

func New(in io.Reader) *Lexer {
	var s scanner.Scanner
	s.Init(in)
	s.Mode ^= scanner.ScanComments // don't skip comments
	s.Whitespace ^= 1 << '\n'      // don't skip tabs and new lines
	l := &Lexer{s: s}
	l.readRune()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var t token.Token

	switch l.ch {
	case '=':
		t = l.either('=', token.EQ, token.ASSIGN)
	case '+':
		t = l.token(token.PLUS)
	case '-':
		t = l.token(token.MINUS)
	case '*':
		t = l.token(token.MUL)
	case '/':
		t = l.token(token.DIV)
	case '!':
		t = l.either('=', token.NOT_EQ, token.BANG)
	case '>':
		t = l.either('=', token.GTE, token.GT)
	case '<':
		t = l.either('=', token.LTE, token.LT)
	case '&':
		t = l.either('&', token.AND, token.VAND)
	case '|':
		p := l.s.Pos()
		switch l.s.Peek() {
		case '|':
			t.Type = token.OR
			t.Literal = "||"
			t.Col = p.Column + 1
			t.Line = p.Line
			l.readRune()
		case '>':
			t.Type = token.PIPE
			t.Literal = "|>"
			t.Col = p.Column + 1
			t.Line = p.Line
			l.readRune()
		default:
			t = l.token(token.VOR)
		}
	case ';':
		t = l.token(token.SEMICOLON)
	case ',':
		t = l.token(token.COMMA)
	case '(':
		t = l.token(token.LPAREN)
	case ')':
		t = l.token(token.RPAREN)
	case '[':
		t = l.token(token.LBRACKET)
	case ']':
		t = l.token(token.RBRACKET)
	case '{':
		t = l.token(token.LBRACE)
	case '}':
		t = l.token(token.RBRACE)
	case '\n':
		t = l.token(token.NEWLINE)
	case scanner.Ident:
		p := l.s.Pos()
		lit := l.s.TokenText()
		col := p.Column
		t = token.Token{
			Type:    token.GetIdentifierType(lit),
			Literal: lit,
			Line:    p.Line,
			Col:     col,
		}
	case scanner.Int, scanner.Float:
		p := l.s.Pos()
		lit := l.s.TokenText()
		t = token.Token{
			Type:    token.NUMBER,
			Literal: lit,
			Line:    p.Line,
			Col:     p.Column,
		}
	case scanner.String:
		p := l.s.Pos()
		lit := l.s.TokenText()
		t = token.Token{
			Type:    token.STRING,
			Literal: lit[1 : len(lit)-1],
			Line:    p.Line,
			Col:     p.Column - 2,
		}
	case scanner.EOF:
		p := l.s.Pos()
		t = token.Token{Type: token.EOF, Literal: "", Line: p.Line, Col: p.Column}
	default:
		p := l.s.Pos()
		lit := l.s.TokenText()
		t = token.Token{Type: token.ERROR, Literal: lit, Line: p.Line, Col: p.Column}
	}

	l.readRune()
	return t
}

func (l *Lexer) readRune() {
	l.ch = l.s.Scan()
}

func (l *Lexer) token(t token.TokenType) token.Token {
	p := l.s.Pos()
	lit := l.s.TokenText()
	return token.Token{Type: t, Literal: lit, Line: p.Line, Col: p.Column}
}

func (l *Lexer) either(what rune, then token.TokenType, otherwise token.TokenType) token.Token {
	p := l.s.Pos()
	lit := l.s.TokenText()
	col := p.Column
	if l.s.Peek() == what {
		l.readRune()
		lit += l.s.TokenText()
		col += 1
		return token.Token{Type: then, Literal: lit, Line: p.Line, Col: col}
	} else {
		return token.Token{Type: otherwise, Literal: lit, Line: p.Line, Col: col}
	}
}
