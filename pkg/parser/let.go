package parser

import (
	"github.com/qiushiyan/peach/pkg/ast"
	"github.com/qiushiyan/peach/pkg/token"
)

func (p *Parser) parseLetStatement() *ast.LetStatement {
	statement := &ast.LetStatement{Token: p.curToken}

	if !p.expectNextToken(token.IDENTIFIER) {
		return nil
	}

	statement.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectNextToken(token.ASSIGN) {
		return nil
	}

	p.advancedToken()
	for !p.endOfExpression() {
		p.advancedToken()
	}

	return statement
}
