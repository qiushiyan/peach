package ast

import "github.com/qiushiyan/qlang/pkg/token"

type ForExpression struct {
	Token      token.Token // the 'for' token
	Iterator   Expression
	IterValues Expression
	Statements *BlockStatement
}

func (fe *ForExpression) expressionNode() {}

func (fe *ForExpression) TokenLiteral() string { return fe.Token.Literal }

func (fe *ForExpression) String() string {
	var out string
	out += fe.TokenLiteral() + " ("
	out += fe.Iterator.String()
	out += " in "
	out += fe.IterValues.String()
	out += ") "
	out += fe.Statements.String()
	return out
}
