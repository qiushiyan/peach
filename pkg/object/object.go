package object

const (
	NUMBER_OBJ   = "NUMBER"
	STRING_OBJ   = "STRING"
	BOOLEAN_OBJ  = "BOOLEAN"
	NULL_OBJ     = "NULL"
	RETURN_OBJ   = "RETURN"
	FUNCTION_OBJ = "FUNCTION"
	ERROR_OBJ    = "ERROR"

	// colors
	COLOR_BLUE  = "\033[34m"
	COLOR_RESET = "\033[0m"
)

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}
