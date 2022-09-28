package ast

import "github.com/qiushiyan/qlang/pkg/token"

// expression node for identifier itself
type Identifier struct {
	Token token.Token // {Type: token.IDENTIFIER, Literal: ?}
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

func (i *Identifier) String() string { return i.Value }
