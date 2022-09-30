package parser

import (
	"github.com/qiushiyan/qlang/pkg/ast"
	"github.com/qiushiyan/qlang/pkg/token"
)

func (p *Parser) parseRangePrefixExpression() ast.Expression {
	expr := &ast.RangeExpression{Token: p.curToken, Start: nil}
	p.advanceToken()
	expr.End = p.parseExpression(LOWEST)
	return expr
}

func (p *Parser) parseRangeInfixExpression(left ast.Expression) ast.Expression {
	expr := &ast.RangeExpression{Token: p.curToken, Start: left}

	if p.nextTokenIs(token.RBRACKET) {
		expr.End = nil
		return expr
	}

	if p.endOfExpression() {
		p.advanceToken()
		p.advanceToken()
		expr.End = nil
		return expr
	}

	p.advanceToken()
	expr.End = p.parseExpression(LOWEST)
	return expr
}
