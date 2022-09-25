package parser

import (
	"strings"
	"testing"

	"github.com/qiushiyan/peach/pkg/ast"
	"github.com/qiushiyan/peach/pkg/lexer"
)

func TestNumberLiteralExpression(t *testing.T) {
	input := "5;"

	l := lexer.New(strings.NewReader(input))
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d",
			len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
			program.Statements[0])
	}

	literal, ok := stmt.Expression.(*ast.NumberLiteral)
	if !ok {
		t.Fatalf("exp not *ast.NumberLiteral. got=%T", stmt.Expression)
	}
	if literal.Value != 5.0 {
		t.Errorf("literal.Value not %v. got=%v", 5.0, literal.Value)
	}
	if literal.TokenLiteral() != "5" {
		t.Errorf("literal.TokenLiteral not %s. got=%s", "5",
			literal.TokenLiteral())
	}
}
