package eval

import (
	"github.com/qiushiyan/qlang/pkg/object"
)

func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR_OBJ
	}
	return false
}
