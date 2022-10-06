package parser

import (
	"github.com/qiushiyan/qlang/pkg/ast"
	"github.com/qiushiyan/qlang/pkg/token"
)

var precedences = map[token.TokenType]int{
	token.ASSIGN:       ASSIGN,
	token.ASSIGN_ARROW: ASSIGN,
	token.EQ:           EQUALS,
	token.NOT_EQ:       EQUALS,
	token.AND:          LESSGREATER,
	token.VAND:         LESSGREATER,
	token.OR:           LESSGREATER,
	token.VOR:          LESSGREATER,
	token.LT:           LESSGREATER,
	token.LTE:          LESSGREATER,
	token.GT:           LESSGREATER,
	token.GTE:          LESSGREATER,
	token.PLUS:         SUM,
	token.MINUS:        SUM,
	token.DIV:          PRODUCT,
	token.MOD:          PRODUCT,
	token.MUL:          PRODUCT,
	token.PIPE:         PIPE,
	token.COLON:        RANGE,
	token.LPAREN:       CALL,
	token.LBRACKET:     INDEX,
}

func (p *Parser) registerAllInfixParsers() {
	// artihmetic
	p.registerInfix(token.PLUS, p.parseInfixExpression)
	p.registerInfix(token.MINUS, p.parseInfixExpression)
	p.registerInfix(token.DIV, p.parseInfixExpression)
	p.registerInfix(token.MOD, p.parseInfixExpression)
	p.registerInfix(token.MUL, p.parseInfixExpression)
	// logical
	p.registerInfix(token.EQ, p.parseInfixExpression)
	p.registerInfix(token.NOT_EQ, p.parseInfixExpression)
	p.registerInfix(token.LT, p.parseInfixExpression)
	p.registerInfix(token.LTE, p.parseInfixExpression)
	p.registerInfix(token.GT, p.parseInfixExpression)
	p.registerInfix(token.GTE, p.parseInfixExpression)
	p.registerInfix(token.OR, p.parseInfixExpression)
	p.registerInfix(token.VOR, p.parseInfixExpression)
	p.registerInfix(token.AND, p.parseInfixExpression)
	p.registerInfix(token.VAND, p.parseInfixExpression)
	// assignment
	p.registerInfix(token.ASSIGN, p.parseAssignExpression)
	p.registerInfix(token.ASSIGN_ARROW, p.parseAssignExpression)
	// function call
	p.registerInfix(token.LPAREN, p.parseCallExpression)
	p.registerInfix(token.PIPE, p.parseInfixExpression)
	// index
	p.registerInfix(token.LBRACKET, p.parseIndexExpression)
	// range
	p.registerInfix(token.COLON, p.parseRangeInfixExpression)
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.nextToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expr := &ast.InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left,
	}
	precedence := p.curPrecedence()
	p.advanceToken()
	expr.Right = p.parseExpression(precedence)
	return expr
}
