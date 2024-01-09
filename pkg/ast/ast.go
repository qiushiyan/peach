package ast

import "bytes"

type Node interface {
	TokenLiteral() string
	String() string
}

// Abstract statement node
type Statement interface {
	Node
	statementNode() // dummy method for distinguishing from expression node
}

// Abstract expression node
type Expression interface {
	Node
	expressionNode() // dummy method for distinguishing from statement node
}

// slice of statement nodes
type Program struct {
	Statements []Statement
}

func (p *Program) String() string {
	var out bytes.Buffer

	for i := range p.Statements {
		out.WriteString(p.Statements[i].String())
	}

	return out.String()
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}
