package eval

import (
	"github.com/qiushiyan/qlang/pkg/ast"
	"github.com/qiushiyan/qlang/pkg/object"
)

func evalPrefixExpression(node *ast.PrefixExpression, env *object.Env) object.Object {
	right := Eval(node.Right, env)
	if object.IsError(right) {
		return right
	}
	switch op := node.Operator; op {
	case "!":
		return evalBangOperatorExpression(right)
	case "+":
		return evalPlusPrefixOperatorExpression(right)
	case "-":
		return evalMinusPrefixOperatorExpression(right)
	default:
		return object.NewError("invalid operator %s for type %s", op, right.Type())
	}
}

func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case object.TRUE:
		return object.FALSE
	case object.FALSE:
		return object.TRUE
	// !null is still NULL
	case object.NULL:
		return object.NULL
	default:
		// !nonzero number = false
		// !zero number = true
		if number, ok := right.(*object.Number); ok {
			if number.Value != 0 {
				return object.FALSE
			} else {
				return object.TRUE
			}
		}
		return newPrefixError("!", right)
	}
}

func evalPlusPrefixOperatorExpression(right object.Object) object.Object {
	switch right.Type() {
	case object.NUMBER_OBJ:
		return right
	default:
		return newPrefixError("+", right)
	}
}

func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	switch right.Type() {
	case object.NUMBER_OBJ:
		value := right.(*object.Number).Value
		return &object.Number{Value: -value}
	default:
		return newPrefixError("-", right)
	}
}

func newPrefixError(operator string, right object.Object) object.Object {
	return object.NewError("invalid operator %s for type %s", operator, right.Type())
}
