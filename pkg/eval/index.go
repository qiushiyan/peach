package eval

import (
	"github.com/qiushiyan/qlang/pkg/ast"
	"github.com/qiushiyan/qlang/pkg/object"
)

func evalIndexExpression(node *ast.IndexExpression, env *object.Env) object.Object {
	left := Eval(node.Left, env)
	if object.IsError(left) {
		return left
	}
	index := Eval(node.Index, env)
	if object.IsError(index) {
		return index
	}
	switch l := left.(type) {
	case object.IVector:
		switch i := index.(type) {
		case *object.Number, *object.Range:
			return evalVectorIndexExpression(l, i)
		default:
			return object.NewError("index must be a number or range, got %s", i.Type())
		}
	case *object.Dict:
		switch index.(type) {
		case object.Hashable:
			return evalDictIndexExpression(left, index)
		default:
			return object.NewError("index must be a hashable type")
		}
	default:
		return object.NewError("invliad index operation for type %s", left.Type())
	}
}

func evalVectorIndexExpression(vec object.IVector, index object.Object) object.Object {
	switch index := index.(type) {
	case *object.Number:
		if idx := int(index.Value) - 1; idx < vec.Length() && idx >= 0 {
			return vec.Values()[idx]
		} else {
			return object.NewError("index out of bounds for vector of length %d, got %d", vec.Length(), int(index.Value))
		}
	case *object.Range:
		start, end, valid := getIndexBounds(index, vec.Length())
		if !valid {
			return object.NewError("index out of bounds for vector of length %d, got %d:%d", vec.Length(), index.Start, index.End)
		}
		return vec.Slice(start, end)
	default:
		return object.NewError("index must be a number or range, got %s", index.Type())
	}
}

// start, end, valid
func getIndexBounds(index *object.Range, length int) (int, int, bool) {
	var start, end int
	// :end
	if index.Start == -1 {
		start = 0
	} else {
		start = index.Start - 1
	}
	// start:
	if index.End == -1 {
		end = length
	} else {
		end = index.End
	}
	if indexOutofBounds(start, end, length) || start > end {
		return 0, 0, false
	}

	return start, end, true
}

func indexOutofBounds(start, end, length int) bool {
	return start < 0 || end > length
}

// index dict
func evalDictIndexExpression(d object.Object, index object.Object) object.Object {
	// index is already verified as Hashable in evalIndexExpression
	dict := d.(*object.Dict)
	if val, ok := dict.Get(index); ok {
		return val
	} else {
		return object.NULL
	}
}
