package parser

import (
	"strings"
	"testing"

	"github.com/qiushiyan/qlang/pkg/ast"
	"github.com/qiushiyan/qlang/pkg/lexer"
)

func TestParsingIndexExpressions(t *testing.T) {
	input := "myVector[1 + 1]"
	l := lexer.New(strings.NewReader(input))
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	statement, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	expr, ok := statement.Expression.(*ast.IndexExpression)
	if !ok {
		t.Fatalf("exp not *ast.IndexExpression. got=%T", statement.Expression)
	}
	if !testIdentifier(t, expr.Left, "myVector") {
		return
	}
	if !testInfixExpression(t, expr.Index, 1, "+", 1) {
		return
	}
}
