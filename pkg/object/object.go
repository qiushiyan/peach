package object

const (
	NUMBER_OBJ = "NUMBER"
	STRING_OBJ = "STRING"
	BOOL_OBJ   = "BOOL"
	NULL_OBJ   = "NULL"
)

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}
