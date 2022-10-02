package parser

import (
	"github.com/qiushiyan/qlang/pkg/ast"
)

func (p *Parser) parseAssignExpression(left ast.Expression) ast.Expression {
	// curToken is token.ASSIGN
	expr := &ast.AssignExpression{Token: p.curToken, Left: left}
	p.advanceToken()
	expr.Value = p.parseExpression(LOWEST)
	return expr
}
