package parser

import (
	"github.com/qiushiyan/qlang/pkg/ast"
	"github.com/qiushiyan/qlang/pkg/token"
)

func (p *Parser) parseAssignStatement() ast.Statement {
	// curToken is the identifier and next token is token.ASSIGN
	statement := &ast.LetStatement{
		Token: token.Token{
			Type:    token.LET,
			Literal: "let",
			Line:    p.curToken.Line,
			Col:     p.curToken.Col,
		},
		Name: &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal},
	}
	p.advanceToken()
	p.advanceToken()
	statement.Value = p.parseExpression(LOWEST)
	if p.endOfExpression() {
		p.advanceToken()
	}
	return statement
}
