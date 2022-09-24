package parser

import (
	"fmt"

	"github.com/qiushiyan/peach/pkg/ast"
	"github.com/qiushiyan/peach/pkg/token"
)

func (p *Parser) parseExpressionStatement() ast.Statement {
	statement := &ast.ExpressionStatement{Token: p.curToken}
	statement.Expression = p.parseExpression(LOWEST)
	p.advanceToken()

	if p.endOfExpression() {
		p.advanceToken()
	}

	return statement
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
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

func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}
