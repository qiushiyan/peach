package object

const (
	NUMBER_OBJ  = "NUMBER"
	STRING_OBJ  = "STRING"
	BOOLEAN_OBJ = "BOOLEAN"
	NULL_OBJ    = "NULL"
	RETURN_OBJ  = "RETURN"
	ERROR_OBJ   = "ERROR"

	// colors
	BLUE  = "\033[34m"
	RESET = "\033[0m"
)

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}
