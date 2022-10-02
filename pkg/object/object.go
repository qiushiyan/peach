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
	// dict
	DICT_OBJ = "DICT"

	// colors
	COLOR_BLUE  = "\033[34m"
	COLOR_RESET = "\033[0m"
)

// only need one "constant" object for true, false and null
var NULL = &Null{}
var TRUE = &Boolean{Value: true}
var FALSE = &Boolean{Value: false}

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}
