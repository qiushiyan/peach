package ast

import "github.com/qiushiyan/qlang/pkg/token"

type ForStatement struct {
	Token     token.Token // the 'for' token
	Condition Expression
	Body      *BlockStatement
}
