package lexer

import (
	"testing"

	"github.com/qiushiyan/peach/pkg/token"
)

func TestNextToken(t *testing.T) {
	input := "=+(){},;||>"

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.ILLEGAL, "|"},
		{token.PIPE, "|>"},
		{token.EOF, ""},
	}

	lexer := New(input)

	for i, test := range tests {
		tok := lexer.NextToken()
		if tok.Type != test.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, test.expectedType, tok.Type)
		}

		if tok.Literal != test.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, test.expectedLiteral, tok.Literal)
		}
	}
}
