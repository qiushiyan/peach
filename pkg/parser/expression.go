package parser

import "github.com/qiushiyan/peach/pkg/ast"

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	statement := &ast.ExpressionStatement{Token: p.curToken}
	statement.Expression = p.parseExpression(LOWEST)
	p.advancedToken()

	for !p.endOfExpression() {
		p.advancedToken()
	}

	return statement
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		return nil
	}
	leftExp := prefix()
	return leftExp
}
