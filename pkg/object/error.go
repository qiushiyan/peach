package object

type Error struct {
	Message string
}

func (e *Error) Inspect() string { return COLOR_BLUE + "ERROR: " + e.Message + COLOR_RESET }
func (e *Error) Type() ObjectType {
	return ERROR_OBJ
}
