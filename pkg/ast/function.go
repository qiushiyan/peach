package ast

import (
	"bytes"
	"strings"

	"github.com/qiushiyan/qlang/pkg/token"
)

type FunctionLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Defaults   map[string]Expression
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode()      {}
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer
	params := make([]string, len(fl.Parameters))
	for i := range fl.Parameters {
		params[i] = fl.Parameters[i].String()
	}
	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())
	return out.String()
}
