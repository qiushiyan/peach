package parser

import (
	"fmt"

	"github.com/qiushiyan/qlang/pkg/token"

	"github.com/qiushiyan/qlang/pkg/ast"
)

func (p *Parser) parseFunction() ast.Expression {
	fn := &ast.FunctionLiteral{Token: p.curToken}
	if !p.expectNextToken(token.LPAREN) {
		return nil
	}
	fn.Parameters = p.parseFunctionParameters()
	p.advanceToken()
	if p.curTokenIs(token.LBRACE) {
		fn.Body = p.parseBlockStatement()
	} else {
		fn.Body = p.parseOnelineBlockStatement()
	}

	// make the last line of the function body the return value (if not already a return statement)
	if len(fn.Body.Statements) > 0 {
		lastStatement := fn.Body.Statements[len(fn.Body.Statements)-1]
		if _, ok := lastStatement.(*ast.ReturnStatement); !ok {
			lastExpression, ok := lastStatement.(*ast.ExpressionStatement)
			if ok {
				fn.Body.Statements[len(fn.Body.Statements)-1] = &ast.ReturnStatement{
					Token: token.Token{
						Type:    token.RETURN,
						Literal: "return",
						Line:    lastExpression.Token.Line,
						Col:     lastExpression.Token.Col},
					Value: lastExpression.Expression,
				}
			} else {
				p.errors = append(p.errors, fmt.Sprintf("can't convert last line of function body to return statement %d:%d", p.curToken.Line, p.curToken.Col))
			}

		}
	}

	return fn
}

func (p *Parser) parseFunctionParameters() []*ast.Identifier {
	identifiers := []*ast.Identifier{}
	if p.nextTokenIs(token.RPAREN) {
		p.advanceToken()
		return identifiers
	}
	p.advanceToken()
	ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	identifiers = append(identifiers, ident)
	for p.nextTokenIs(token.COMMA) {
		p.advanceToken()
		p.advanceToken()
		identifier := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
		identifiers = append(identifiers, identifier)
	}
	if !p.expectNextToken(token.RPAREN) {
		return nil
	}
	return identifiers
}
