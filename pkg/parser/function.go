package parser

import (
	"github.com/qiushiyan/peach/pkg/token"

	"github.com/qiushiyan/peach/pkg/ast"
)

func (p *Parser) parseFunction() ast.Expression {
	fn := &ast.FunctionLiteral{Token: p.curToken}
	if !p.expectNextToken(token.LPAREN) {
		return nil
	}
	fn.Parameters = p.parseFunctionParameters()
	if !p.expectNextToken(token.LBRACE) {
		return nil
	}
	fn.Body = p.parseBlockStatement()
	return fn
}

func (p *Parser) parseFunctionParameters() []*ast.Identifier {
	identifiers := []*ast.Identifier{}
	if p.nextTokenIs(token.RPAREN) {
		p.advanceToken()
		return identifiers
	}
	p.advanceToken()
	ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	identifiers = append(identifiers, ident)
	for p.nextTokenIs(token.COMMA) {
		p.advanceToken()
		p.advanceToken()
		identifier := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
		identifiers = append(identifiers, identifier)
	}
	if !p.expectNextToken(token.RPAREN) {
		return nil
	}
	return identifiers
}
