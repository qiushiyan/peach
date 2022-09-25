package token

const (
	ERROR TokenType = iota
	EOF

	// Identifiers
	IDENTIFIER // variable names: x, y, add, foobar

	// literals
	NUMBER
	STRING

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
	OR     // ||
	VOR    // |
	AND    // &&
	VAND   // &

	// Delimiters
	COMMA     // ,
	SEMICOLON // ;
	LPAREN    // (
	RPAREN    // )
	LBRACKET  // [
	RBRACKET  // ]
	LBRACE    // {
	RBRACE    // }
	NEWLINE   // \n

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
	Col     int
	Line    int
}

// Helper to create a Token, Literal can be a byte or a string
func New[T byte | []byte](t TokenType, l T) Token {
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
	case EOF:
		return "EOF"
	case IDENTIFIER:
		return "IDENTIFIER"
	case NUMBER:
		return "NUMBER"
	case STRING:
		return "STRING"
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
	case OR:
		return "OR"
	case VOR:
		return "VOR"
	case AND:
		return "AND"
	case VAND:
		return "VAND"
	case COMMA:
		return "COMMA"
	case SEMICOLON:
		return "SEMICOLON"
	case LPAREN:
		return "LPAREN"
	case RPAREN:
		return "RPAREN"
	case LBRACKET:
		return "LBRACKET"
	case RBRACKET:
		return "RBRACKET"
	case LBRACE:
		return "LBRACE"
	case RBRACE:
		return "RBRACE"
	case NEWLINE:
		return "NEWLINE"
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
