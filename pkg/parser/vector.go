package parser

import (
	"github.com/qiushiyan/qlang/pkg/ast"
	"github.com/qiushiyan/qlang/pkg/token"
)

func (p *Parser) parseVectorLiteral() ast.Expression {
	Vector := &ast.VectorLiteral{Token: p.curToken}

	Vector.Elements = p.parseExpressionList(token.RBRACKET)

	return Vector
}
