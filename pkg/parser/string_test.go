package parser

import (
	"strings"
	"testing"

	"github.com/qiushiyan/peach/pkg/ast"
	"github.com/qiushiyan/peach/pkg/lexer"
)

func TestStringLiteralExpression(t *testing.T) {
	input := "\"hello world\";"

	l := lexer.New(strings.NewReader(input))
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d",
			len(program.Statements))
	}
	statement, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
			program.Statements[0])
	}

	literal, ok := statement.Expression.(*ast.StringLiteral)
	if !ok {
		t.Fatalf("exp not *ast.StringLiteral. got=%T", statement.Expression)
	}
	if literal.Value != "hello world" {
		t.Errorf("literal.Value not %s. got=%s", "hello world", literal.Value)
	}
	if literal.TokenLiteral() != "hello world" {
		t.Errorf("literal.TokenLiteral not %s. got=%s", "5",
			literal.TokenLiteral())
	}
}
