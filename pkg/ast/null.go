package ast

import "github.com/qiushiyan/peach/pkg/token"

type Null struct {
	Token token.Token
	Value bool
}

func (n *Null) expressionNode()      {}
func (n *Null) TokenLiteral() string { return n.Token.Literal }
func (n *Null) String() string       { return n.Token.Literal }
