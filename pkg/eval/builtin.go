package eval

import (
	"fmt"

	"github.com/qiushiyan/qlang/pkg/object"
)

var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			case *object.Number:
				return &object.Integer{Value: 1}
			case *object.Vector:
				return &object.Integer{Value: int64(arg.Length())}
			case *object.NumericVector:
				return &object.Integer{Value: int64(arg.Length())}
			case *object.CharacterVector:
				return &object.Integer{Value: int64(arg.Length())}
			case *object.LogicalVector:
				return &object.Integer{Value: int64(arg.Length())}
			default:
				return newError("invalid type %s for function `len()`", args[0].Type())
			}
		},
		ParametersNum: 1,
	},
	"head": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) == 0 {
				return newError("head() requires at least 1 argument")
			}
			switch arg := args[0].(type) {
			// case *object.Vector:
			// 	return arg.Head()
			// case *object.NumericVector:
			// 	return arg.Head()
			// case *object.CharacterVector:
			// 	return arg.Head()
			// case *object.LogicalVector:
			// 	return arg.Head()
			default:
				fmt.Println(arg)
				return newError("invalid type %s for function `head()`", args[0].Type())
			}
		},
		ParametersNum: 1,
	},
}
