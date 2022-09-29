package eval

import "github.com/qiushiyan/qlang/pkg/object"

var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			return NULL
		}},
}
