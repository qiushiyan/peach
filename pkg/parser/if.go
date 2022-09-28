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
	expr.Consequence = p.parseIfBranch()
	if p.nextTokenIs(token.ELSE) {
		p.advanceToken()
		expr.Alternative = p.parseIfBranch()
	}

	return expr
}

func (p *Parser) parseIfBranch() *ast.BlockStatement {
	// jump to the next token after ) or ELSE
	p.advanceToken()
	if p.curTokenIs(token.LBRACE) {
		return p.parseBlockStatement()
	} else {
		return &ast.BlockStatement{
			Token:      token.Token{Type: token.LBRACE, Literal: "{", Col: p.curToken.Col, Line: p.curToken.Line},
			Statements: []ast.Statement{p.parseExpressionStatement()},
		}
	}
}
