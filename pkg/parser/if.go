package parser

import (
	"github.com/qiushiyan/qlang/pkg/ast"
	"github.com/qiushiyan/qlang/pkg/token"
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

	p.advanceToken()
	// curToken is the one after ) in if condition
	expr.Consequence = p.parseIfBranch()
	// curToken is }
	p.advanceToken()
	if p.curTokenIs(token.ELSE) {
		p.advanceToken()
		expr.Alternative = p.parseIfBranch()
	}

	return expr
}

func (p *Parser) parseIfBranch() *ast.BlockStatement {
	if p.curTokenIs(token.LBRACE) {
		return p.parseBlockStatement()
	} else {
		return p.parseOnelineBlockStatement()
	}
}
