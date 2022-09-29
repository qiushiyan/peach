package parser

import (
	"github.com/qiushiyan/qlang/pkg/ast"
)

// return either of
// *ast.Identifier
// *ast.LetStatement
func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}
