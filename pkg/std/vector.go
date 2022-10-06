package std

import (
	"github.com/qiushiyan/qlang/pkg/object"
)

func vectorCreate(env *object.Env, args ...object.Object) object.Object {
	num, ok := args[0].(*object.Number)
	if !ok {
		return object.NewError("invalid first argument in function `vector()`, must be a non-negative number, got %s", num.Type())
	}
	if num.Value < 0 {
		return object.NewError("invalid first argument in function `vector()`, must be a non-negative number, got %f", num.Value)
	}

	vals := make([]object.Object, int(num.Value))
	for i := 0; i < int(num.Value); i++ {
		vals[i] = object.NULL
	}
	return &object.Vector{BaseVector: object.BaseVector{Elements: vals}}
}

func vectorHead(env *object.Env, args ...object.Object) object.Object {
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
}

func vectorTail(env *object.Env, args ...object.Object) object.Object {
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
}

func vectorAppend(env *object.Env, args ...object.Object) object.Object {
	switch arg := args[0].(type) {
	case object.IVector:
		return arg.Append(args[1:]...)
	default:
		return object.NewError("invalid first argument in `append()`, must be a vector, got %s", args[0].Type())
	}
}

func vectorConvert(env *object.Env, args ...object.Object) object.Object {
	switch arg := args[0].(type) {
	case object.IVector:
		return arg
	case *object.Range:
		if arg.Start == -1 || arg.End == -1 {
			return object.NewError("invalid range in `as_vector()`, must be a finite range")
		}
		results := make([]object.Object, arg.End-arg.Start+1)
		for i := arg.Start; i <= arg.End; i++ {
			results[i-arg.Start] = &object.Number{Value: float64(i)}
		}
		return object.NewVector(results)
	default:
		return object.NewError("invalid first argument in `convert()`, must be either one of vector, range or dict, got %s", args[0].Type())
	}
}
