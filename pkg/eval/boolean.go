package eval

import "github.com/qiushiyan/qlang/pkg/object"

func evalBoolean(value bool) *object.Boolean {
	if value {
		return object.TRUE
	} else {
		return object.FALSE
	}
}
