package ast

import (
	"bytes"
	"strings"

	"github.com/qiushiyan/qlang/pkg/token"
)

type VectorLiteral struct {
	Token    token.Token // the '[' token
	Elements []Expression
}

func (al *VectorLiteral) expressionNode()      {}
func (al *VectorLiteral) TokenLiteral() string { return al.Token.Literal }
func (al *VectorLiteral) String() string {
	var out bytes.Buffer
	elements := make([]string, len(al.Elements))
	for i := range al.Elements {
		elements[i] = al.Elements[i].String()
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	return out.String()
}
