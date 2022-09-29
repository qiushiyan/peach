package parser

import (
	"github.com/qiushiyan/qlang/pkg/ast"
	"github.com/qiushiyan/qlang/pkg/token"
)

func (p *Parser) parseIndexExpression(left ast.Expression) ast.Expression {
	expr := &ast.IndexExpression{Token: p.curToken, Left: left}
	p.advanceToken()
	expr.Index = p.parseExpression(LOWEST)
	if !p.expectNextToken(token.RBRACKET) {
		return nil
	}
	return expr
}
