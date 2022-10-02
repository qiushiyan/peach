package std

import (
	"fmt"

	"github.com/qiushiyan/qlang/pkg/object"
)

func length(env *object.Env, args ...object.Object) object.Object {
	switch arg := args[0].(type) {
	case *object.String:
		return &object.Number{Value: float64(len(arg.Value))}
	case *object.Number:
		return &object.Number{Value: 1}
	case object.IVector:
		return &object.Number{Value: float64(arg.Length())}
	case *object.Dict:
		return &object.Number{Value: float64(len(arg.Pairs))}
	default:
		return object.NewError("invalid first argument in function `len()`, can be a number, string, or vector, got %s", args[0].Type())
	}
}

func print(env *object.Env, args ...object.Object) object.Object {
	for _, arg := range args {
		fmt.Println(arg.Inspect())
	}
	return object.NULL
}
