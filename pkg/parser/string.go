package parser

import "github.com/qiushiyan/peach/pkg/ast"

func (p *Parser) parseStringLiteral() ast.Expression {
	return &ast.StringLiteral{Token: p.curToken, Value: p.curToken.Literal}
}
