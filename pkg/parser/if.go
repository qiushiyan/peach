package parser

import (
	"github.com/qiushiyan/peach/pkg/ast"
	"github.com/qiushiyan/peach/pkg/token"
)

func (p *Parser) parseIfExpression() ast.Expression {
	expr := &ast.IfExpression{Token: p.curToken}
	if !p.expectNextToken(token.LPAREN) {
		return nil
	}
	p.advanceToken()
	expr.Condition = p.parseExpression(LOWEST)
	if !p.expectNextToken(token.RPAREN) {
		return nil
	}
	if !p.expectNextToken(token.LBRACE) {
		return nil
	}
	expr.Consequence = p.parseBlockStatement()

	if p.nextTokenIs(token.ELSE) {
		p.advanceToken()
		if !p.expectNextToken(token.LBRACE) {
			return nil
		}
		expr.Alternative = p.parseBlockStatement()
	}

	return expr
}
