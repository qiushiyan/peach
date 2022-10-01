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
	fn.Parameters, fn.Defaults = p.parseFunctionParameters()
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

func (p *Parser) parseFunctionParameters() ([]*ast.Identifier, map[string]ast.Expression) {
	parameters := []*ast.Identifier{}
	defaults := make(map[string]ast.Expression)

	// curToken is LPAREN
	if p.nextTokenIs(token.RPAREN) {
		p.advanceToken()
		return parameters, defaults
	}

	for !p.nextTokenIs(token.RPAREN) {
		p.advanceToken()
		if p.curTokenIs(token.EOF) {
			p.errors = append(p.errors, fmt.Sprintf("unterminated function parameters %d:%d", p.curToken.Line, p.curToken.Col))
			return nil, nil
		}

		expr := p.parseExpression(LOWEST)
		if identifier, ok := expr.(*ast.Identifier); ok {
			parameters = append(parameters, identifier)
		}
		if assignExpr, ok := expr.(*ast.AssignExpression); ok {
			parameters = append(parameters, assignExpr.Name)
			defaults[assignExpr.Name.Value] = assignExpr.Value
		}

		if p.nextTokenIs(token.COMMA) {
			p.advanceToken()
		}
	}

	if !p.expectNextToken(token.RPAREN) {
		return nil, nil
	}
	return parameters, defaults
}
