package eval

import (
	"github.com/qiushiyan/qlang/pkg/ast"
	"github.com/qiushiyan/qlang/pkg/object"
)

func evalForExpression(node *ast.ForExpression, env *object.Env) object.Object {
	var iterValues object.Object
	switch iter := node.IterValues.(type) {
	case *ast.RangeExpression:
		result := CheckRange(iter, env, true)
		if object.IsError(result) {
			return result
		}
		iterValues = result
	default:
		iterValues = Eval(node.IterValues, env)
	}
	if object.IsError(iterValues) {
		return iterValues
	}

	identifier, ok := node.Iterator.(*ast.Identifier)
	if !ok {
		return object.NewError("iterator should be an identifier, got %s", node.Iterator.String())
	}

	switch it := iterValues.(type) {
	case object.IVector:
		return evalForVectorExpression(node, env, it, identifier)
	case *object.Dict:
		return evalForDictExpression(node, env, it, identifier)
	case *object.Range:
		return evalForRangeExpression(node, env, it, identifier)
	default:
		return &object.Error{Message: "for loops can only operate on vector or dict"}
	}
}

func evalForVectorExpression(node *ast.ForExpression, env *object.Env, iterValues object.IVector, iterator *ast.Identifier) object.Object {
	var result object.Object
	forEnv := object.NewEnclosedEnvironment(env)
	for _, val := range iterValues.Values() {
		forEnv.Set(iterator.Value, val)
		result = evalBlockStatement(node.Statements, forEnv)
		if object.IsError(result) {
			return result
		}
	}
	return result
}

func evalForDictExpression(node *ast.ForExpression, env *object.Env, iterValues *object.Dict, iterator *ast.Identifier) object.Object {
	var result object.Object
	for _, key := range iterValues.Keys() {
		forEnv := object.NewEnclosedEnvironment(env)
		forEnv.Set(iterator.Value, key)
		result = evalBlockStatement(node.Statements, forEnv)
		if object.IsError(result) {
			return result
		}
	}
	return result
}

func evalForRangeExpression(node *ast.ForExpression, env *object.Env, iterValues *object.Range, iterator *ast.Identifier) object.Object {
	var result object.Object
	forEnv := object.NewEnclosedEnvironment(env)
	for i := iterValues.Start; i < iterValues.End; i++ {
		forEnv.Set(iterator.Value, &object.Number{Value: float64(i)})
		result = evalBlockStatement(node.Statements, forEnv)
		if object.IsError(result) {
			return result
		}
	}
	return result
}
