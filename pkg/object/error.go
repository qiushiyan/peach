package object

type Error struct {
	Message string
}

func (e *Error) Inspect() string { return BLUE + "ERROR: " + e.Message + RESET }
func (e *Error) Type() ObjectType {
	return ERROR_OBJ
}
