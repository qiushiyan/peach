package parser

import (
	"fmt"

	"github.com/qiushiyan/qlang/pkg/ast"
	"github.com/qiushiyan/qlang/pkg/token"
)

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{Token: p.curToken}
	block.Statements = []ast.Statement{}

	p.advanceToken()
	for !p.curTokenIs(token.RBRACE) && !p.curTokenIs(token.EOF) {
		statement := p.parseStatement()
		if statement != nil {
			block.Statements = append(block.Statements, statement)
		}
		p.advanceToken()
	}
	if !p.curTokenIs(token.RBRACE) {
		p.errors = append(p.errors, fmt.Sprintf("unclosed block, missing '}' at %d:%d", p.curToken.Line, p.curToken.Col))
		return nil
	}
	return block
}
