package eval

import "github.com/qiushiyan/qlang/pkg/object"

var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			case *object.Number:
				return &object.Integer{Value: 1}
			default:
				return newError("invalid type %s for function `len()`", args[0].Type())
			}
		},
		ParametersNum: 1,
	},
}
