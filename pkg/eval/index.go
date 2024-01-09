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
		case *object.Number, *object.Range, object.IVector:
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
	case object.IVector:
		return evalVectorIndexVectorExpression(vec, index)
	default:
		return object.NewError("index must be one of number, range and vector, got %s", index.Type())
	}
}

func evalVectorIndexVectorExpression(vec object.IVector, index object.IVector) object.Object {
	if index.Length() == 0 {
		return &object.Error{Message: "index vector can not be empty"}
	}
	if index.Length() > vec.Length() {
		return object.NewError("index vector length can not be greater than vector length, got %d > %d", index.Length(), vec.Length())
	}
	// var values []object.Object
	values := make([]object.Object, 0, index.Length())
	switch index := index.(type) {
	case *object.NumericVector:
		for i := range index.Values() {
			el := index.Values()[i]
			idx := int(el.(*object.Number).Value)
			if idx := idx - 1; idx < vec.Length() && idx >= 0 {
				values[i] = vec.Values()[idx]
			} else {
				return object.NewError("index out of bounds for vector of length %d, got %d", vec.Length(), idx)
			}
		}
	case *object.LogicalVector:
		if index.Length() != vec.Length() {
			return object.NewError("boolean indexing must use a vector of same size, got %d != %d", index.Length(), vec.Length())
		}
		for idx, val := range vec.Values() {
			b := index.Values()[idx].(*object.Boolean).Value
			if b {
				values = append(values, val)
			}
		}
	default:
		return object.NewError("index vector must be a numerical or logical vector, got %s", index.ElementType())
	}
	return object.NewVector(values)
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
