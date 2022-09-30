package object

const (
	NUMBER_OBJ   = "NUMBER"
	STRING_OBJ   = "STRING"
	BOOLEAN_OBJ  = "BOOLEAN"
	NULL_OBJ     = "NULL"
	RETURN_OBJ   = "RETURN"
	FUNCTION_OBJ = "FUNCTION"
	BUILTIN_OBJ  = "BUILTIN"
	ERROR_OBJ    = "ERROR"
	RANGE_OBJ    = "RANGE"
	// vector
	VECTOR_OBJ = "VECTOR"
	// []object.Number
	NUMERIC_VECTOR_OBJ = "NUMERIC_VECTOR"
	// []object.String
	CHARACTER_VECTOR_OBJ = "CHARACTER_VECTOR"
	// []object.Boolean
	LOGICAL_VECTOR_OBJ = "LOGICAL_VECTOR"

	// dict
	DICT_OBJ = "DICT"

	// colors
	COLOR_BLUE  = "\033[34m"
	COLOR_RESET = "\033[0m"
)

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}
