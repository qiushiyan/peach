package object

const (
	NUMBER_OBJ  = "NUMBER"
	STRING_OBJ  = "STRING"
	BOOLEAN_OBJ = "BOOLEAN"
	NULL_OBJ    = "NULL"
	RETURN_OBJ  = "RETURN"
)

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}
