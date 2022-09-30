package eval

import (
	"github.com/qiushiyan/qlang/pkg/ast"
	"github.com/qiushiyan/qlang/pkg/object"
)

func evalCallExpression(node *ast.CallExpression, env *object.Env) object.Object {
	fn := Eval(node.Function, env)
	if object.IsError(fn) {
		return fn
	}

	args := evalExpressions(node.Arguments, env)
	if len(args) == 1 && object.IsError(args[0]) {
		return args[0]
	}

	if identifier, ok := node.Function.(*ast.Identifier); ok {
		return applyFunction(fn, args, identifier.Value)
	} else {
		return applyFunction(fn, args, nil)
	}

}

func applyFunction(fn object.Object, args []object.Object, name interface{}) object.Object {
	var fnName string
	if name != nil {
		fnName = name.(string)
	} else {
		fnName = fn.Inspect()
	}

	switch fn := fn.(type) {
	case *object.Function:
		if err := checkParameters(fnName, len(args), fn.RequiredParametersNum); err != nil {
			return err
		}
		fnEnv := makeFunctionEnv(fn, args)
		evaluated := Eval(fn.Body, fnEnv)
		return unwrapReturnValue(evaluated)
	case *object.Builtin:
		if err := checkParameters(fnName, len(args), fn.RequiredParametersNum); err != nil {
			return err
		}
		return fn.Fn(args...)
	default:
		return object.NewError("%s is not a function", fnName)
	}
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

func checkParameters(fnName string, argLength int, paramLength int) object.Object {
	if argLength < paramLength {
		return object.NewError("not enough arguments for %s, got=%d, required=%d", fnName, argLength, paramLength)
	}
	return nil
}
