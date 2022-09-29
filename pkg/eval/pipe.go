package eval

import (
	"github.com/qiushiyan/qlang/pkg/ast"
	"github.com/qiushiyan/qlang/pkg/object"
	"github.com/qiushiyan/qlang/pkg/token"
)

func evalPipeExpression(node *ast.InfixExpression, env *object.Env) *ast.CallExpression {
	if call, isCall := node.Right.(*ast.CallExpression); isCall {
		call.Arguments = append([]ast.Expression{node.Left}, call.Arguments...)
		return call
	}

	if identifier, isIdentifier := node.Right.(*ast.Identifier); isIdentifier {
		function := Eval(identifier, env)
		if function == nil || isError(function) {
			return nil
		}
		call := &ast.CallExpression{
			Token: token.Token{
				Type:    token.LPAREN,
				Literal: "(",
				Col:     node.Token.Col,
				Line:    node.Token.Line,
			},
			Function:  node.Right,
			Arguments: []ast.Expression{node.Left},
		}
		return call
	}

	return nil
}
