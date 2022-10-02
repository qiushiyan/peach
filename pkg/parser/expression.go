package parser

import (
	"fmt"

	"github.com/qiushiyan/qlang/pkg/ast"
	"github.com/qiushiyan/qlang/pkg/token"
)

func (p *Parser) parseExpressionStatement() ast.Statement {
	statement := &ast.ExpressionStatement{Token: p.curToken}
	statement.Expression = p.parseExpression(LOWEST)

	if p.endOfExpression() {
		p.advanceToken()
	}
	return statement
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	if p.curTokenIs(token.NEWLINE) {
		p.advanceToken()
	}
	prefixFn := p.prefixParseFns[p.curToken.Type]
	if prefixFn == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}
	leftExp := prefixFn()

	for !p.endOfExpression() && precedence < p.peekPrecedence() {
		infixFn := p.infixParseFns[p.nextToken.Type]
		if infixFn == nil {
			return leftExp
		}
		p.advanceToken()
		leftExp = infixFn(leftExp)
	}

	return leftExp
}

// parse a comma separated list of expressions into a slice
func (p *Parser) parseExpressionList(end token.TokenType) []ast.Expression {
	list := []ast.Expression{}
	if p.nextTokenIs(end) {
		p.advanceToken()
		return list
	}
	p.advanceToken()
	list = append(list, p.parseExpression(LOWEST))
	for p.nextTokenIs(token.COMMA) {
		p.advanceToken()
		// curToken is now comma
		// allow multiple newlines between comma
		for p.nextTokenIs(token.NEWLINE) {
			p.advanceToken()
		}
		p.advanceToken()
		// curToken is now the next expression
		// allow trailing commas
		if p.curTokenIs(end) {
			return list
		}
		list = append(list, p.parseExpression(LOWEST))
	}
	for p.nextTokenIs(token.NEWLINE) {
		p.advanceToken()
	}
	if !p.expectNextToken(end) {
		return nil
	}
	return list
}

func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}
