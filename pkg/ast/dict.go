package ast

import "github.com/qiushiyan/qlang/pkg/token"

type DictLiteral struct {
	Token token.Token // the '{' token
	Pairs map[Expression]Expression
}

func (ml *DictLiteral) expressionNode() {}

func (ml *DictLiteral) TokenLiteral() string { return ml.Token.Literal }

func (ml *DictLiteral) String() string {
	var out string
	out += "{\n"
	for k, v := range ml.Pairs {
		out += k.String()
		out += ": "
		out += v.String()
		out += ", "
		out += "\n"
	}
	out += "\n}"
	return out
}
