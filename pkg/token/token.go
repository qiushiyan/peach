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
	COLON     // :
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
	NULL     // null
	RETURN   // return
	FOR      // for
	IN       // IN
)

type TokenType byte

type Token struct {
	Type    TokenType
	Literal string
	Col     int
	Line    int
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
	"null":   NULL,
	"return": RETURN,
	"for":    FOR,
	"in":     IN,
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
	case COLON:
		return "COLON"
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
	case NULL:
		return "NULL"
	case RETURN:
		return "RETURN"
	case FOR:
		return "FOR"
	case IN:
		return "IN"
	}
	return "UNKNOWN"
}
