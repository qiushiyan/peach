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
			return newError("range start must be a number")
		}
	}
	if node.End == nil {
		end = -1
	} else {
		end_ := Eval(node.End, env)
		if end_, ok := end_.(*object.Number); ok {
			end = int(end_.Value)
		} else {
			return newError("range end must be a number")
		}
	}
	return &object.Range{Start: start, End: end}
}
