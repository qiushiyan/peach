package eval

import (
	"github.com/qiushiyan/qlang/pkg/ast"
	"github.com/qiushiyan/qlang/pkg/object"
)

func evalDictLiteral(node *ast.DictLiteral, env *object.Env) object.Object {
	pairs := make(map[object.DictKey]object.DictPair)
	for keyNode, valueNode := range node.Pairs {
		var key object.Object
		if identifier, ok := keyNode.(*ast.Identifier); ok {
			if val, ok := env.Get(identifier.Value); ok {
				key = val
			} else {
				key = &object.String{Value: identifier.String()}
			}
		} else {
			key = Eval(keyNode, env)
		}

		if object.IsError(key) {
			return key
		}
		dictKey, ok := key.(object.Hashable)
		if !ok {
			return object.NewError("unusable as dict key: %s", key.Type())
		}

		value := Eval(valueNode, env)
		if object.IsError(value) {
			return value
		}
		hashed := dictKey.Hash()
		pairs[hashed] = object.DictPair{Key: key, Value: value}
	}
	return &object.Dict{Pairs: pairs}
}
