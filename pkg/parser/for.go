package parser

import (
	"github.com/qiushiyan/qlang/pkg/ast"
	"github.com/qiushiyan/qlang/pkg/token"
)

func (p *Parser) parseForExpression() ast.Expression {
	expr := &ast.ForExpression{Token: p.curToken}
	if !p.expectNextToken(token.LPAREN) {
		return nil
	}
	p.advanceToken()
	expr.Iterator = p.parseExpression(LOWEST)
	if !p.expectNextToken(token.IN) {
		return nil
	}
	p.advanceToken()
	expr.IterValues = p.parseExpression(LOWEST)
	if !p.expectNextToken(token.RPAREN) {
		return nil
	}

	p.advanceToken()
	// curToken is the one after ) in for condition
	if p.curTokenIs(token.LBRACE) {
		expr.Statements = p.parseBlockStatement()
	} else {
		expr.Statements = p.parseOnelineBlockStatement()
	}

	return expr
}
