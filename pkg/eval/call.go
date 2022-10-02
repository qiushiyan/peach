package eval

import (
	"github.com/qiushiyan/qlang/pkg/ast"
	"github.com/qiushiyan/qlang/pkg/object"
)

func evalCallExpression(node *ast.CallExpression, env *object.Env) object.Object {
	// node.Function: object.Function
	/// node.Arguments: explicit arguments
	fn := Eval(node.Function, env)
	if object.IsError(fn) {
		return fn
	}

	args, err, ok := evalArguments(node.Arguments, env)
	if err != nil {
		return err
	}
	if !ok {
		// at this point argumenta are evaluated successfully, but at least one kwarg appears before args
		return object.NewError("Keyword arguments must appear after positional arguments")
	}

	if identifier, ok := node.Function.(*ast.Identifier); ok {
		return applyFunction(fn, args, identifier.Value, env)
	} else {
		return applyFunction(fn, args, nil, env)
	}

}

type Arguments struct {
	Args   []object.Object
	Kwargs map[string]object.Object
}

// reuturn 3 values
// the parsed map, whether parsing is successful, error object
func evalArguments(args []ast.Expression, env *object.Env) (*Arguments, object.Object, bool) {
	result := &Arguments{Args: []object.Object{}, Kwargs: map[string]object.Object{}}
	for _, expr := range args {
		if assignExpr, ok := expr.(*ast.AssignExpression); ok {
			identifier, ok := assignExpr.Left.(*ast.Identifier)
			if !ok {
				return nil, object.NewError("invalid keyword arguments, left hand side must be an identifier, got %s", assignExpr.Left.String()), false
			}
			val := Eval(assignExpr.Value, env)
			result.Kwargs[identifier.Value] = val
			if object.IsError(val) {
				return result, val, false
			}
		} else {
			// only allow kwargs after args
			if len(result.Kwargs) > 0 {
				return nil, nil, false
			}
			val := Eval(expr, env)
			result.Args = append(result.Args, val)
			if object.IsError(val) {
				return result, val, false
			}
		}
	}

	return result, nil, true
}

func applyFunction(fn object.Object, args *Arguments, name interface{}, env *object.Env) object.Object {
	var fnName string
	if name != nil {
		fnName = name.(string)
	} else {
		fnName = fn.Inspect()
	}

	argLength := len(args.Args) + len(args.Kwargs)
	switch fn := fn.(type) {
	case *object.Function:
		if err := checkParameters(fnName, argLength, fn.RequiredParametersNum); err != nil {
			return err
		}
		fnEnv, errorKwarg := makeFunctionEnv(fn, args, argLength)
		if errorKwarg != "" {
			return object.NewError("unexpected keyword argument %s in function %s", errorKwarg, fnName)
		}
		evaluated := Eval(fn.Body, fnEnv)
		return unwrapReturnValue(evaluated)
	case *object.Builtin:
		if err := checkParameters(fnName, argLength, fn.RequiredParametersNum); err != nil {
			return err
		}
		if len(args.Kwargs) > 0 {
			return object.NewError("builtin function %s does not support keyword arguments", fnName)
		}
		return fn.Fn(env, args.Args...)
	default:
		return object.NewError("%s is not a function", fnName)
	}
}

func makeFunctionEnv(fn *object.Function, args *Arguments, argLength int) (*object.Env, string) {
	env := object.NewEnclosedEnvironment(fn.Env)
	// set default values
	for param, val := range fn.Defaults {
		env.Set(param, Eval(val, env))
	}

	for idx, positionalArg := range args.Args {
		env.Set(fn.Parameters[idx].Value, positionalArg)
	}
	for name, kwarg := range args.Kwargs {
		kwargFound := false
		// if kwarg is found in function parameters, set it
		for _, param := range fn.Parameters {
			if param.Value == name {
				env.Set(name, kwarg)
				kwargFound = true
				break
			}
		}
		if !kwargFound {
			return env, name
		}
	}
	return env, ""
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
