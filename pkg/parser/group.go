package parser

import (
	"github.com/qiushiyan/peach/pkg/token"

	"github.com/qiushiyan/peach/pkg/ast"
)

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.advanceToken()
	exp := p.parseExpression(LOWEST)

	// this should be a right parenthesis because it has lowest precedence
	if !p.expectNextToken(token.RPAREN) {
		return nil
	}
	return exp
}
