package std

import "github.com/qiushiyan/qlang/pkg/object"

func dictKeys(env *object.Env, args ...object.Object) object.Object {
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
}

func dictValues(env *object.Env, args ...object.Object) object.Object {
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
}
