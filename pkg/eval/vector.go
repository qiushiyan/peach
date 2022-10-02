package eval

import (
	"github.com/qiushiyan/qlang/pkg/ast"
	"github.com/qiushiyan/qlang/pkg/object"
)

func evalVectorLiteral(node ast.Node, env *object.Env) object.Object {
	vectorNode := node.(*ast.VectorLiteral)
	elements := evalExpressions(vectorNode.Elements, env)

	firstElement := elements[0]
	if len(elements) == 1 && object.IsError(firstElement) {
		return elements[0]
	}

	return object.NewVector(elements)
}

func evalVectorInfixExpression(operator string, left, right object.IVector) object.Object {
	return left.Infix(func(x object.Object, y object.Object) object.Object {
		return evalInfixExpression(operator, x, y)
	}, right)
}
