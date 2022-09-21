package ast

import (
	"testing"

	"github.com/qiushiyan/peach/pkg/token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "LET"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENTIFIER, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENTIFIER, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	stringExpected := "LET myVar = anotherVar;"
	if program.String() != stringExpected {
		t.Errorf("program.String() wrong. expect=%q got=%q", stringExpected, program.String())
	}
}
