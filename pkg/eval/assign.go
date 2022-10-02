package eval

import (
	"github.com/qiushiyan/qlang/pkg/ast"
	"github.com/qiushiyan/qlang/pkg/object"
)

func evalAssignExpression(node *ast.AssignExpression, env *object.Env) object.Object {
	switch left := node.Left.(type) {
	case *ast.Identifier:
		val := Eval(node.Value, env)
		if object.IsError(val) {
			return val
		}
		env.Set(left.Value, val)
		return object.NULL

	case *ast.IndexExpression:
		replacement := Eval(node.Value, env)
		if object.IsError(replacement) {
			return replacement
		}
		// the object to be indexed
		obj := Eval(left.Left, env)
		index := Eval(left.Index, env)
		switch obj := obj.(type) {
		case object.IVector:
			return evalAssignVectorIndexExpression(obj, index, replacement)
		case *object.Dict:
			return evalAssignDictIndexExpression(obj, index, replacement)
		default:
			return object.NewError("index operator not supported: %s", obj.Type())
		}
	default:
		return object.NewError("invlalid assignment, must be an identifier or index expression got %s", left.String())
	}
}

func evalAssignVectorIndexExpression(vec object.IVector, index object.Object, replacement object.Object) object.Object {
	switch index := index.(type) {
	case *object.Number:
		idx := int(index.Value) - 1
		if idx < vec.Length() && idx >= 0 {
			if err := vec.Set(idx, replacement); object.IsError(err) {
				return err
			}
		} else {
			return object.NewError("index out of bounds for vector of length %d, got %d", vec.Length(), int(index.Value))
		}
		return object.NULL
	case *object.Range:
		start, end, valid := getIndexBounds(index, vec.Length())
		if !valid {
			return object.NewError("index out of bounds for vector of length %d, got %d:%d", vec.Length(), index.Start, index.End)
		}
		switch newvals := replacement.(type) {
		case *object.Number, *object.String, *object.Boolean:
			for i := start; i < end; i++ {
				vec.Set(i, newvals)
			}
		case object.IVector:
			if (end - start) != newvals.Length() {
				return object.NewError("length of replacement must match length of range, got %d, expected %d", newvals.Length(), end)
			}
			if err := vec.Replace(start, end, newvals); object.IsError(err) {
				return err
			}
		default:
			return object.NewError("invalid values to replace for vector index assignment, got %s", newvals.Type())
		}
	default:
		return object.NewError("index must be a number or range")
	}
	return object.NULL
}

func evalAssignDictIndexExpression(dict *object.Dict, index object.Object, replacement object.Object) object.Object {
	dict.Set(index, replacement)
	return object.NULL
}
