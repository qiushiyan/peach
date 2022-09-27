package parser

import "github.com/qiushiyan/peach/pkg/ast"

func (p *Parser) parseNull() ast.Expression {
	return &ast.Null{}
}
