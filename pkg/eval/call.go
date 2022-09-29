package eval

import (
	"github.com/qiushiyan/qlang/pkg/ast"
	"github.com/qiushiyan/qlang/pkg/object"
)

func evalCallExpression(node *ast.CallExpression, env *object.Env) object.Object {
	fn := Eval(node.Function, env)
	if isError(fn) {
		return fn
	}

	args := evalParameters(node.Arguments, env)
	if len(args) == 1 && isError(args[0]) {
		return args[0]
	}

	if identifier, ok := node.Function.(*ast.Identifier); ok {
		return applyFunction(fn, args, identifier.Value)
	} else {
		return applyFunction(fn, args, nil)
	}

}

func applyFunction(fn object.Object, args []object.Object, name interface{}) object.Object {
	switch fn := fn.(type) {
	case *object.Function:
		fnEnv := makeFunctionEnv(fn, args)
		evaluated := Eval(fn.Body, fnEnv)
		return unwrapReturnValue(evaluated)
	case *object.Builtin:
		return fn.Fn(args...)
	default:
		if name != nil {
			return newError("%s is not a function", name)
		} else {
			return newError("not a function: %s", fn.Type())
		}
	}
}

func evalParameters(exps []ast.Expression, env *object.Env) []object.Object {
	var result []object.Object
	for _, e := range exps {
		evaluated := Eval(e, env)
		if isError(evaluated) {
			return []object.Object{evaluated}
		}
		result = append(result, evaluated)
	}
	return result
}

func makeFunctionEnv(fn *object.Function, args []object.Object) *object.Env {
	env := object.NewEnclosedEnvironment(fn.Env)
	for paramIdx, param := range fn.Parameters {
		env.Set(param.Value, args[paramIdx])
	}
	return env
}

func unwrapReturnValue(obj object.Object) object.Object {
	if returnValue, ok := obj.(*object.ReturnValue); ok {
		return returnValue.Value
	}
	return obj
}
