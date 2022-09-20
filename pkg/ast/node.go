package ast

import "github.com/qiushiyan/peach/pkg/token"

// Statement node
type Statement interface {
	Node
	statementNode() // dummy method for distinguishing from expression node
}

// Expression node
type Expression interface {
	Node
	expressionNode() // dummy method for distinguishing from statement node
}

// statement node for `let <identifier> = <expression>`
type LetStatement struct {
	Token token.Token // {Type: token.LET, Literal: "LET"}
	Name  *Identifier
	Value Expression
}

func NewLetStatement(Name *Identifier, Value Expression) LetStatement {
	return LetStatement{
		Token: token.Token{Type: token.LET, Literal: "LET"},
		Name:  Name,
		Value: Value,
	}
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// statement node for `return <expression>`
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func NewReturnStatement(ReturnValue Expression) ReturnStatement {
	return ReturnStatement{
		Token:       token.Token{Type: token.RETURN, Literal: "RETURN"},
		ReturnValue: ReturnValue,
	}
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

// expression node for identifier itself
type Identifier struct {
	Token token.Token // {Type: token.IDENTIFIER, Literal: ?}
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
