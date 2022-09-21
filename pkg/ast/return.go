package ast

import (
	"bytes"

	"github.com/qiushiyan/peach/pkg/token"
)

// statement node for `return <expression>`
type ReturnStatement struct {
	Token token.Token
	Value Expression
}

func NewReturnStatement(ReturnValue Expression) ReturnStatement {
	return ReturnStatement{
		Token: token.Token{Type: token.RETURN, Literal: "RETURN"},
		Value: ReturnValue,
	}
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.Value != nil {
		out.WriteString(rs.Value.String())
	}
	out.WriteString(";")
	return out.String()
}
