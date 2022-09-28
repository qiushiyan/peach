package parser

import "github.com/qiushiyan/qlang/pkg/ast"

func (p *Parser) parseNull() ast.Expression {
	return &ast.Null{}
}
