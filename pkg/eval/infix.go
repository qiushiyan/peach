package eval

import "github.com/qiushiyan/qlang/pkg/object"

func evalInfixExpression(operator string, left object.Object, right object.Object) object.Object {
	switch {
	case left.Type() == object.NUMBER_OBJ && right.Type() == object.NUMBER_OBJ:
		return evalNumberInfixExpression(operator, left, right)
	case left.Type() == object.STRING_OBJ && right.Type() == object.STRING_OBJ:
		return evalStringInfixExpression(operator, left, right)
	case operator == "==":
		return evalBoolean(left == right)
	case operator == "!=":
		return evalBoolean(left != right)
	case left.Type() != right.Type():
		return object.NewError("operaor %s needs values of the same type, got %s %s %s", operator, left.Type(), operator, right.Type())
	default:
		return newInfixError(left, operator, right)
	}
}

func newInfixError(left object.Object, operator string, right object.Object) object.Object {
	return object.NewError("invalid operation %s %s %s", left.Type(), operator, right.Type())
}
