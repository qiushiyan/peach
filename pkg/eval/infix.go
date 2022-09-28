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
		return newError("operaor %s needs values of the same type, got %s and %s", operator, left.Type(), right.Type())
	default:
		return newError("operator %s does not support %s and %s", operator, left.Type(), right.Type())
	}
}
