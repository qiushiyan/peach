package token

type TokenType byte

const (
	ILLEGAL TokenType = iota
	EOF

	// Identifiers + literals
	IDENTIFIER // variable names: x, y, add, foobar
	INT

	// Operators
	ASSIGN
	PLUS

	// Delimiters
	COMMA
	SEMICOLON
	LPAREN
	RPAREN
	LBRACE
	RBRACE

	// Keywords or special contructs
	FUNCTION
	LET
	PIPE
)

func New(t TokenType, l byte) Token {
	return Token{Type: t, Literal: string(l)}
}

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func GetIdentifierType(k string) TokenType {
	if tokenType, ok := keywords[k]; ok {
		return tokenType
	}
	return IDENTIFIER
}

type Token struct {
	Type    TokenType
	Literal string
}
