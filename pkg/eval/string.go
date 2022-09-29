package eval

import "github.com/qiushiyan/qlang/pkg/object"

func evalStringLiteral(value string) object.Object {
	return &object.String{Value: value}
}

func evalStringInfixExpression(operator string, left object.Object, right object.Object) object.Object {
	leftVal := left.(*object.String).Value
	rightVal := right.(*object.String).Value

	switch operator {
	case "+":
		return &object.String{Value: leftVal + rightVal}
	default:
		return newInfixError(left, operator, right)
	}
}
