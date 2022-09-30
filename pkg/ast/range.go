package ast

import (
	"bytes"

	"github.com/qiushiyan/qlang/pkg/token"
)

type RangeExpression struct {
	Token token.Token // The token.COLON token
	Start Expression
	End   Expression
}

func (re *RangeExpression) expressionNode()      {}
func (re *RangeExpression) TokenLiteral() string { return re.Token.Literal }
func (re *RangeExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(re.Start.String())
	out.WriteString(":")
	out.WriteString(re.End.String())
	out.WriteString(")")
	return out.String()
}
