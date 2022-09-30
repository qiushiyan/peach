package object

import "fmt"

type Error struct {
	Message string
}

func (e *Error) Inspect() string { return COLOR_BLUE + "ERROR: " + e.Message + COLOR_RESET }
func (e *Error) Type() ObjectType {
	return ERROR_OBJ
}

func NewError(text string, a ...interface{}) *Error {
	return &Error{
		Message: fmt.Sprintf(text, a...),
	}
}
