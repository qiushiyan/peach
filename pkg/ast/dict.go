package ast

import (
	"bytes"

	"github.com/qiushiyan/qlang/pkg/token"
)

type DictLiteral struct {
	Token token.Token // the '{' token
	Pairs map[Expression]Expression
}

func (ml *DictLiteral) expressionNode() {}

func (ml *DictLiteral) TokenLiteral() string { return ml.Token.Literal }

func (ml *DictLiteral) String() string {
	var out bytes.Buffer
	out.WriteString("{\n")
	for k := range ml.Pairs {
		v := ml.Pairs[k]

		out.WriteString(k.String())
		out.WriteString(": ")
		out.WriteString(v.String())
		out.WriteString(", n")
	}
	out.WriteString("\n}")
	return out.String()
}
