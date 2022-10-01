package parser

import (
	"fmt"

	"github.com/qiushiyan/qlang/pkg/ast"
	"github.com/qiushiyan/qlang/pkg/token"
)

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	// curToken is {
	p.deregisterPrefix(token.LBRACE)
	block := &ast.BlockStatement{Token: p.curToken, Statements: []ast.Statement{}}
	p.advanceToken()

	for !p.curTokenIs(token.RBRACE) && !p.curTokenIs(token.EOF) {
		statement := p.parseStatement()
		if statement != nil {
			block.Statements = append(block.Statements, statement)
		}
		p.advanceToken()
		if p.curTokenIs(token.NEWLINE) {
			p.advanceToken()
		}
	}

	if !p.curTokenIs(token.RBRACE) {
		p.unclosedBlockError()
		return nil
	}

	return block
}

// parse a one-liner block statement without {}
// used by function and if/else expressions
// fn(x) x + 1
// if(x > 0) x + 1 else x - 1
func (p *Parser) parseOnelineBlockStatement() *ast.BlockStatement {
	if p.curTokenIs(token.NEWLINE) || p.curTokenIs(token.SEMICOLON) || p.curTokenIs(token.EOF) {
		p.errors = append(p.errors, fmt.Sprintf("incomplete statement at %d:%d", p.curToken.Line, p.curToken.Col))
		return nil
	}

	return &ast.BlockStatement{
		Token: token.Token{
			Type:    token.LBRACE,
			Literal: "{",
			Line:    p.curToken.Line,
			Col:     p.curToken.Col,
		},
		Statements: []ast.Statement{p.parseStatement()},
	}
}

func (p *Parser) unclosedBlockError() {
	p.errors = append(p.errors, fmt.Sprintf("unclosed block, missing '}' at %d:%d", p.curToken.Line, p.curToken.Col))
}
