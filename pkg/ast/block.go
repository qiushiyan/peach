package ast

import (
	"bytes"

	"github.com/qiushiyan/qlang/pkg/token"
)

type BlockStatement struct {
	Token      token.Token // the { token
	Statements []Statement
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer
	for i := range bs.Statements {
		out.WriteString(bs.Statements[i].String())
	}
	return out.String()
}
