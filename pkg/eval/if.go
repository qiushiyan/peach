package eval

import (
	"github.com/qiushiyan/qlang/pkg/ast"
	"github.com/qiushiyan/qlang/pkg/object"
)

func evalIfExpression(ie *ast.IfExpression, env *object.Env) object.Object {
	condition := Eval(ie.Condition, env)
	if object.IsError(condition) {
		return condition
	}

	ifEnv := object.NewEnclosedEnvironment(env)
	if isTruthy(condition) {
		return Eval(ie.Consequence, ifEnv)
	} else if ie.Alternative != nil {
		return Eval(ie.Alternative, ifEnv)
	} else {
		return object.NULL
	}
}

func isTruthy(obj object.Object) bool {
	switch {
	case obj.Type() == object.NULL_OBJ:
		return false
	case obj.Type() == object.BOOLEAN_OBJ:
		if obj.(*object.Boolean).Value {
			return true
		} else {
			return false
		}
	case obj.Type() == object.NUMBER_OBJ:
		if obj.(*object.Number).Value != 0 {
			return true
		} else {
			return false
		}
	default:
		return true
	}
}
