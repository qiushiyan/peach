package ast

import (
	"bytes"

	"github.com/qiushiyan/qlang/pkg/token"
)

// expression node for `<identifier> = <expression>`
type AssignExpression struct {
	Token token.Token // the identifier
	Name  *Identifier
	Value Expression
}

func (ae *AssignExpression) expressionNode()      {}
func (ae *AssignExpression) TokenLiteral() string { return ae.Token.Literal }

func (ae *AssignExpression) String() string {
	var out bytes.Buffer
	out.WriteString(ae.TokenLiteral() + " ")
	out.WriteString(ae.Name.String())
	out.WriteString(" = ")
	if ae.Value != nil {
		out.WriteString(ae.Value.String())
	}
	return out.String()
}
