package eval

import (
	"github.com/qiushiyan/qlang/pkg/ast"
	"github.com/qiushiyan/qlang/pkg/object"
)

func evalForExpression(node *ast.ForExpression, env *object.Env) object.Object {
	var iterValues object.Object
	switch iter := node.IterValues.(type) {
	case *ast.RangeExpression:
		var start, end float64
		if iterStart, ok := iter.Start.(*ast.NumberLiteral); ok {
			if iterStart.Value == -1 {
				return &object.Error{Message: "range start and end must be specified in for loop"}
			}
			start = iterStart.Value
		} else {
			iterStart := Eval(iter.Start, env)
			if object.IsError(iterStart) {
				return iterStart
			}
			if iterStart.Type() != object.NUMBER_OBJ {
				return &object.Error{Message: "range start and end must be number in for loop"}
			}
			start = iterStart.(*object.Number).Value
		}
		if iterEnd, ok := iter.End.(*ast.NumberLiteral); ok {
			if iterEnd.Value == -1 {
				return &object.Error{Message: "range start and end must be specified in for loop"}
			}
			end = iterEnd.Value
		} else {
			iterEnd := Eval(iter.End, env)
			if object.IsError(iterEnd) {
				return iterEnd
			}
			if iterEnd.Type() != object.NUMBER_OBJ {
				return &object.Error{Message: "range start and end must be number in for loop"}
			}
			end = iterEnd.(*object.Number).Value
		}
		iterValues = &object.Range{Start: int(start), End: int(end) + 1} // +1 here for 1-based indexing
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
