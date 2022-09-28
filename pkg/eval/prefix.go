package eval

import "github.com/qiushiyan/qlang/pkg/object"

func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	// !null is still null
	case NULL:
		return NULL
	default:
		// !nonzero number = false
		// !zero number = true
		if number, ok := right.(*object.Number); ok {
			if number.Value != 0 {
				return FALSE
			} else {
				return TRUE
			}
		}
		return FALSE
	}
}

func evalPlusPrefixOperatorExpression(right object.Object) object.Object {
	switch right.Type() {
	case object.NUMBER_OBJ:
		return right
	default:
		return newError("operator + does not support type %s", right.Type())
	}
}

func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	switch right.Type() {
	case object.NUMBER_OBJ:
		value := right.(*object.Number).Value
		return &object.Number{Value: -value}
	default:
		return newError("operator - does not support type %s", right.Type())
	}
}
