package token

const (
	ILLEGAL TokenType = iota
	EOF

	// Identifiers
	IDENTIFIER // variable names: x, y, add, foobar

	// literals
	INTEGER

	// Operators
	ASSIGN // =
	PIPE   // |>
	PLUS   // +
	MINUS  // -
	DIV    // /
	BANG   // !
	MUL    // *
	GT     // >
	LT     // <
	GTE    // >=
	LTE    // <=
	EQ     // ==
	NOT_EQ // !=

	// Delimiters
	COMMA     // ,
	SEMICOLON // ;
	LPAREN    // (
	RPAREN    // )
	LBRACE    // {
	RBRACE    // }

	// Keywords
	FUNCTION // fn
	LET      // let
	IF       // if
	ELSE     // else
	TRUE     // true
	FALSE    // false
	RETURN   // return
	FOR      // for
)

type TokenType byte

type Token struct {
	Type    TokenType
	Literal string
}

func New(t TokenType, l byte) Token {
	return Token{Type: t, Literal: string(l)}
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
	"return": RETURN,
	"for":    FOR,
}

func GetIdentifierType(k string) TokenType {
	if tokenType, ok := keywords[k]; ok {
		return tokenType
	}
	return IDENTIFIER
}

// for debugging purposes
func (t TokenType) String() string {
	switch t {
	case ILLEGAL:
		return "ILLEGAL"
	case EOF:
		return "EOF"
	case IDENTIFIER:
		return "IDENTIFIER"
	case INTEGER:
		return "INTEGER"
	case ASSIGN:
		return "ASSIGN"
	case PIPE:
		return "PIPE"
	case PLUS:
		return "PLUS"
	case MINUS:
		return "MINUS"
	case DIV:
		return "DIV"
	case BANG:
		return "BANG"
	case MUL:
		return "MUL"
	case GT:
		return "GT"
	case LT:
		return "LT"
	case GTE:
		return "GTE"
	case LTE:
		return "LTE"
	case EQ:
		return "EQ"
	case NOT_EQ:
		return "NOT_EQ"
	case COMMA:
		return "COMMA"
	case SEMICOLON:
		return "SEMICOLON"
	case LPAREN:
		return "LPAREN"
	case RPAREN:
		return "RPAREN"
	case LBRACE:
		return "LBRACE"
	case RBRACE:
		return "RBRACE"
	case FUNCTION:
		return "FUNCTION"
	case LET:
		return "LET"
	case IF:
		return "IF"
	case ELSE:
		return "ELSE"
	case TRUE:
		return "TRUE"
	case FALSE:
		return "FALSE"
	case RETURN:
		return "RETURN"
	case FOR:
		return "FOR"
	}
	return "UNKNOWN"
}
