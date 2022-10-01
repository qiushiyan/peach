package parser

import (
	"github.com/qiushiyan/qlang/pkg/ast"
)

func (p *Parser) parseAssignExpression(left ast.Expression) ast.Expression {
	// curToken is the identifier and next token is token.ASSIGN
	if identifier, ok := left.(*ast.Identifier); !ok {
		p.errors = append(p.errors, "invalid assignment, left hand side of = must be an identifier")
		return nil
	} else {
		expr := &ast.AssignExpression{Token: p.curToken, Name: identifier}
		p.advanceToken()
		expr.Value = p.parseExpression(LOWEST)
		return expr
	}
}
