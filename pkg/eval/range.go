package eval

import (
	"github.com/qiushiyan/qlang/pkg/ast"
	"github.com/qiushiyan/qlang/pkg/object"
)

func evalRangeExpression(node *ast.RangeExpression, env *object.Env) object.Object {
	var start, end int
	if node.Start == nil {
		start = -1
	} else {
		start_ := Eval(node.Start, env)
		if start_, ok := start_.(*object.Number); ok {
			start = int(start_.Value)
		} else {
			return object.NewError("range start must be a number")
		}
	}
	if node.End == nil {
		end = -1
	} else {
		end_ := Eval(node.End, env)
		if end_, ok := end_.(*object.Number); ok {
			end = int(end_.Value)
		} else {
			return object.NewError("range end must be a number")
		}
	}
	return &object.Range{Start: start, End: end}
}

// validate a range expression and return the range object
func CheckRange(node *ast.RangeExpression, env *object.Env, isBounded bool) object.Object {
	var iterStart, iterEnd float64
	start := checkRangeEnd(node.Start, env, isBounded)
	switch start := start.(type) {
	case *object.Error:
		return start
	case float64:
		iterStart = start
	}
	end := checkRangeEnd(node.End, env, isBounded)
	switch end := end.(type) {
	case *object.Error:
		return end
	case float64:
		iterEnd = end
	}

	return &object.Range{Start: int(iterStart), End: int(iterEnd) + 1} // +1 here for 1-based indexing
}

// validate range.Start and range.End
// return a go float or object.Error
func checkRangeEnd(expr ast.Expression, env *object.Env, isBounded bool) interface{} {
	var num float64
	if x, ok := expr.(*ast.NumberLiteral); ok {
		if isBounded {
			if x.Value == -1 {
				return &object.Error{Message: "range start and end must be specified in for loop"}
			}
		}
		num = x.Value
	} else {
		x := Eval(expr, env)
		if object.IsError(x) {
			return x
		}
		if x.Type() != object.NUMBER_OBJ {
			return &object.Error{Message: "range start and end must be number in for loop"}
		}
		num = x.(*object.Number).Value
	}
	return num
}
