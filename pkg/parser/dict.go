package parser

import (
	"github.com/qiushiyan/qlang/pkg/ast"
	"github.com/qiushiyan/qlang/pkg/token"
)

func (p *Parser) parseDictLiteral() ast.Expression {
	// curToken is token.LBRACE
	dict := &ast.DictLiteral{Token: p.curToken, Pairs: make(map[ast.Expression]ast.Expression)}

	for !p.nextTokenIs(token.RBRACE) {
		p.advanceToken()

		// disable range parsing for map key
		p.deregisterInfix(token.COLON)
		key := p.parseExpression(LOWEST)
		if !p.expectNextToken(token.COLON) {
			return nil
		}
		p.advanceToken()

		// enable range parsing for map value
		p.registerInfix(token.COLON, p.parseRangeInfixExpression)
		value := p.parseExpression(LOWEST)
		dict.Pairs[key] = value

		for p.nextTokenIs(token.NEWLINE) {
			p.advanceToken()
		}
		if !p.nextTokenIs(token.RBRACE) && !p.expectNextToken(token.COMMA) {
			return nil
		}
	}
	if !p.expectNextToken(token.RBRACE) {
		return nil
	}

	p.registerInfix(token.COLON, p.parseRangeInfixExpression)
	return dict
}
