package parser

import "github.com/qiushiyan/peach/pkg/ast"

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	statement := &ast.ReturnStatement{Token: p.curToken}
	p.advancedToken()

	for !p.endOfExpression() {
		p.advancedToken()
	}

	return statement
}
