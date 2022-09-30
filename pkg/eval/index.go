package eval

import (
	"github.com/qiushiyan/qlang/pkg/ast"
	"github.com/qiushiyan/qlang/pkg/object"
)

func evalIndexExpression(node *ast.IndexExpression, env *object.Env) object.Object {
	left := Eval(node.Left, env)
	if isError(left) {
		return left
	}
	index := Eval(node.Index, env)
	if isError(index) {
		return index
	}
	switch left.(type) {
	case *object.Vector, *object.NumericVector, *object.CharacterVector, *object.LogicalVector:
		return evalVectorIndexExpression(left, index)
	default:
		return newError("invliad index operation for type %s", left.Type())
	}
}

func evalVectorIndexExpression(v object.Object, index object.Object) object.Object {
	switch v.Type() {
	case object.VECTOR_OBJ:
		if start, end, valid := getIndexBounds(index, len(v.(*object.Vector).Elements)); valid {
			if end == 0 {
				return v.(*object.Vector).Elements[start]
			} else {
				return &object.Vector{Elements: v.(*object.Vector).Elements[start:end]}
			}
		} else {
			return newError("index out of bounds for vector")
		}
	case object.NUMERIC_VECTOR_OBJ:
		if start, end, valid := getIndexBounds(index, len(v.(*object.NumericVector).Elements)); valid {
			if end == 0 {
				return v.(*object.NumericVector).Elements[start]
			} else {
				return &object.NumericVector{Elements: v.(*object.NumericVector).Elements[start:end]}
			}
		} else {
			return newError("index out of bounds for numeric vector")
		}
	case object.CHARACTER_VECTOR_OBJ:
		if start, end, valid := getIndexBounds(index, len(v.(*object.CharacterVector).Elements)); valid {
			if end == 0 {
				return v.(*object.CharacterVector).Elements[start]
			} else {
				return &object.CharacterVector{Elements: v.(*object.CharacterVector).Elements[start:end]}
			}
		} else {
			return newError("index out of bounds for character vector")
		}
	case object.LOGICAL_VECTOR_OBJ:
		if start, end, valid := getIndexBounds(index, len(v.(*object.LogicalVector).Elements)); valid {
			if end == 0 {
				return v.(*object.LogicalVector).Elements[start]
			} else {
				return &object.LogicalVector{Elements: v.(*object.LogicalVector).Elements[start:end]}
			}
		} else {
			return newError("index out of bounds for logical vector")
		}
	default:
		return newError("invliad index operation for type %s", v.Type())
	}
}

// start, end, valid
func getIndexBounds(index object.Object, length int) (int, int, bool) {
	switch index.(type) {
	case *object.Number:
		idx := int(index.(*object.Number).Value) - 1
		if idx < 0 || idx >= length {
			return 0, 0, false
		}
		return idx, 0, true
	case *object.Range:
		var start, end int
		rangeObj := index.(*object.Range)
		if rangeObj.Start == -1 {
			start = 0
		} else {
			start = rangeObj.Start - 1
		}
		if rangeObj.End == -1 {
			end = length
		} else {
			end = rangeObj.End
		}
		if indexOutofBounds(start, end, length) || start > end {
			return 0, 0, false
		}

		return start, end, true
	default:
		return 0, 0, false
	}
}

func indexOutofBounds(start, end, length int) bool {
	return start < 0 || end > length
}
