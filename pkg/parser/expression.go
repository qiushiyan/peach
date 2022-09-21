package parser

import "github.com/qiushiyan/peach/pkg/ast"

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	statement := &ast.ExpressionStatement{Token: p.curToken}

	p.advancedToken()

	for !p.endOfExpression() {
		p.advancedToken()
	}

	return statement
}
