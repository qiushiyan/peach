package parser

import (
	"github.com/qiushiyan/peach/pkg/ast"
	"github.com/qiushiyan/peach/pkg/token"
)

func (p *Parser) parseCallExpression(function ast.Expression) ast.Expression {
	expr := &ast.CallExpression{Token: p.curToken, Function: function}
	expr.Arguments = p.parseCallArguments()
	return expr
}

func (p *Parser) parseCallArguments() []ast.Expression {
	var args []ast.Expression
	if p.nextTokenIs(token.RPAREN) {
		p.advanceToken()
		return args
	}
	p.advanceToken()
	args = append(args, p.parseExpression(LOWEST))
	for p.nextToken.Type == token.COMMA {
		p.advanceToken()
		p.advanceToken()
		args = append(args, p.parseExpression(LOWEST))
	}
	if !p.expectNextToken(token.RPAREN) {
		return nil
	}
	return args
}
