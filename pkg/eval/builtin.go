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
			case object.IVector:
				return &object.Integer{Value: int64(arg.Length())}
			default:
				return object.NewError("invalid first argument in function `len()`, can be a number, string, or vector, got %s", args[0].Type())
			}
		},
		RequiredParametersNum: 1,
	},
	"head": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			var n int
			if len(args) == 1 {
				n = 6
			} else {
				num, ok := args[1].(*object.Number)
				if !ok {
					return &object.Error{Message: "invalid second argument in function `head()`, must be a non-negative number"}
				}
				n = int(num.Value)
			}

			switch arg := args[0].(type) {
			case object.IVector:
				if n < 0 {
					return &object.Error{Message: "invalid second argument in `head()`, must be a non-negative number"}
				}
				return arg.Head(n)
			default:
				return object.NewError("invalid first argument in `head()`, must be a vector, got %s", args[0].Type())
			}
		},
		RequiredParametersNum: 1,
	},
	"tail": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			var n int
			if len(args) == 1 {
				n = 6
			} else {
				num, ok := args[1].(*object.Number)
				if !ok {
					return &object.Error{Message: "invalid second argument in function `tail()`, must be a non-negative number"}
				}
				n = int(num.Value)
			}

			switch arg := args[0].(type) {
			case object.IVector:
				if n < 0 {
					return &object.Error{Message: "invalid second argument in `head()`, must be a non-negative number"}
				}
				return arg.Tail(n)
			default:
				return object.NewError("invalid first argument in `tail()`, must be a vector, got %s", args[0].Type())
			}
		},
		RequiredParametersNum: 1,
	},
	"append": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			switch arg := args[0].(type) {
			case object.IVector:
				return arg.Append(args[1:]...)
			default:
				return object.NewError("invalid first argument in `append()`, must be a vector, got %s", args[0].Type())
			}
		},
		RequiredParametersNum: 2,
	},
	"keys": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			switch arg := args[0].(type) {
			case *object.Dict:
				var keys []object.Object
				for _, pair := range arg.Pairs {
					keys = append(keys, pair.Key)
				}
				return object.NewVector(keys)
			default:
				return object.NewError("invalid first argument in `keys()`, must be a dict, got %s", args[0].Type())
			}
		},
		RequiredParametersNum: 1,
	},
	"values": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			switch arg := args[0].(type) {
			case *object.Dict:
				var values []object.Object
				for _, pair := range arg.Pairs {
					values = append(values, pair.Value)
				}
				return object.NewVector(values)
			default:
				return object.NewError("invalid first argument in `values()`, must be a dict, got %s", args[0].Type())
			}
		},
		RequiredParametersNum: 1,
	},
	"print": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}
			return NULL
		},
		RequiredParametersNum: 1,
	},
}
