package ast

import (
	"bytes"

	"github.com/qiushiyan/peach/pkg/token"
)

// statement node for `let <identifier> = <expression>`
type LetStatement struct {
	Token token.Token // {Type: token.LET, Literal: "LET"}
	Name  *Identifier
	Value Expression
}

func NewLetStatement(Name *Identifier, Value Expression) LetStatement {
	return LetStatement{
		Token: token.Token{Type: token.LET, Literal: "LET"},
		Name:  Name,
		Value: Value,
	}
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

func (ls *LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()
}
