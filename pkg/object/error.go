package object

import (
	"fmt"
)

type Error struct {
	Message string
}

func (e *Error) Inspect() string {
	return e.Message
}

func (e *Error) Type() ObjectType {
	return ERROR_OBJ
}

func NewError(text string, a ...interface{}) *Error {
	return &Error{
		Message: fmt.Sprintf(text, a...),
	}
}

func IsError(obj Object) bool {
	if obj != nil {
		return obj.Type() == ERROR_OBJ
	}
	return false
}
