package parser

import (
	"github.com/qiushiyan/peach/pkg/ast"
)

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	statement := &ast.ReturnStatement{Token: p.curToken}
	p.advanceToken()

	statement.Value = p.parseExpression(LOWEST)

	if p.endOfExpression() {
		p.advanceToken()
	}

	return statement
}
