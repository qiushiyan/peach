package ast

import (
	"bytes"
	"strings"

	"github.com/qiushiyan/qlang/pkg/token"
)

type CallExpression struct {
	Token     token.Token // The '(' token
	Function  Expression  // Function literal or identifier
	Arguments []Expression
}

func (ce *CallExpression) expressionNode()      {}
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) String() string {
	var out bytes.Buffer
	args := make([]string, len(ce.Arguments))
	for i := range ce.Arguments {
		args[i] = ce.Arguments[i].String()
	}
	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")
	return out.String()
}
