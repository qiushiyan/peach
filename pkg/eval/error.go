package eval

import (
	"fmt"

	"github.com/qiushiyan/qlang/pkg/object"
)

func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR_OBJ
	}
	return false
}

func newError(text string, a ...interface{}) *object.Error {
	return &object.Error{
		Message: fmt.Sprintf(text, a...),
	}
}
