package eval

import (
	"github.com/qiushiyan/qlang/pkg/object"
)

func evalInfixExpression(operator string, left object.Object, right object.Object) object.Object {
	switch {
	case left.Type() == object.NUMBER_OBJ && right.Type() == object.NUMBER_OBJ:
		return evalNumberInfixExpression(operator, left, right)
	case left.Type() == object.STRING_OBJ && right.Type() == object.STRING_OBJ:
		return evalStringInfixExpression(operator, left, right)
	// the next few cases enable infix operations between vectors and scalars
	case left.Type() == object.NUMBER_OBJ && right.Type() == object.VECTOR_OBJ:
		// left and right are swapped because BaseVector.Infix() checks length for the second argument
		return evalVectorInfixExpression(operator, right.(object.IVector), newLengthOneVector(left))
	case left.Type() == object.VECTOR_OBJ && right.Type() == object.NUMBER_OBJ:
		return evalVectorInfixExpression(operator, left.(object.IVector), newLengthOneVector(right))
	case left.Type() == object.STRING_OBJ && right.Type() == object.VECTOR_OBJ:
		return evalVectorInfixExpression(operator, right.(object.IVector), newLengthOneVector(left))
	case left.Type() == object.VECTOR_OBJ && right.Type() == object.STRING_OBJ:
		return evalVectorInfixExpression(operator, left.(object.IVector), newLengthOneVector(right))
	case left.Type() == object.BOOLEAN_OBJ && right.Type() == object.VECTOR_OBJ:
		return evalVectorInfixExpression(operator, right.(object.IVector), newLengthOneVector(left))
	case left.Type() == object.VECTOR_OBJ && right.Type() == object.BOOLEAN_OBJ:
		return evalVectorInfixExpression(operator, left.(object.IVector), newLengthOneVector(right))
	case operator == "&":
		if left.Type() != object.VECTOR_OBJ || right.Type() != object.VECTOR_OBJ {
			return object.NewError("& should be used for vector comparison, got %s & %s", left.Type(), right.Type())
		}
		return evalVectorInfixExpression("&&", left.(object.IVector), right.(object.IVector))
	case operator == "|":
		if left.Type() != object.VECTOR_OBJ || right.Type() != object.VECTOR_OBJ {
			return object.NewError("| should be used for vector comparison, got %s | %", left.Type(), right.Type())
		}
		return evalVectorInfixExpression("||", left.(object.IVector), right.(object.IVector))
	case left.Type() == object.VECTOR_OBJ && right.Type() == object.VECTOR_OBJ:
		left := left.(object.IVector)
		right := right.(object.IVector)
		return evalVectorInfixExpression(operator, left, right)
	case operator == "==":
		return evalBoolean(left == right)
	case operator == "!=":
		return evalBoolean(left != right)
	case operator == "&&":
		return evalBoolean(left == object.TRUE && right == object.TRUE)
	case operator == "||":
		return evalBoolean(left == object.TRUE || right == object.TRUE)
	case left.Type() != right.Type():
		return object.NewError("operaor %s needs values of the same type, got %s %s %s", operator, left.Type(), operator, right.Type())
	default:
		return newInfixError(left, operator, right)
	}
}

func newInfixError(left object.Object, operator string, right object.Object) object.Object {
	return object.NewError("invalid operation %s %s %s", left.Type(), operator, right.Type())
}

func newLengthOneVector(obj object.Object) object.IVector {
	switch obj.(type) {
	case *object.Number:
		return &object.NumericVector{BaseVector: object.BaseVector{Elements: []object.Object{obj}}}
	case *object.String:
		return &object.CharacterVector{BaseVector: object.BaseVector{Elements: []object.Object{obj}}}
	case *object.Boolean:
		return &object.LogicalVector{BaseVector: object.BaseVector{Elements: []object.Object{obj}}}
	default:
		return &object.Vector{BaseVector: object.BaseVector{Elements: []object.Object{obj}}}
	}
}
