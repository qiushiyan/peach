package parser

import (
	"strings"
	"testing"

	"github.com/qiushiyan/qlang/pkg/ast"
	"github.com/qiushiyan/qlang/pkg/lexer"
)

func TestAssign(t *testing.T) {
	tests := []struct {
		input              string
		expectedIdentifier string
		expectedValue      interface{}
	}{
		{"x = 5", "x", 5.0},
		{"y = true;", "y", true},
		{"foobar = y;", "foobar", "y"},
	}

	for _, tt := range tests {
		l := lexer.New(strings.NewReader(tt.input))
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain 1 statements. got=%d",
				len(program.Statements))
		}

		statement := program.Statements[0].(*ast.ExpressionStatement)

		expr, ok := statement.Expression.(*ast.AssignExpression)
		if !ok {
			t.Fatalf("statement.Expression is not ast.AssignExpression. got=%T",
				statement.Expression)
		}

		testIdentifier(t, expr.Name, tt.expectedIdentifier)
		testLiteralExpression(t, expr.Value, tt.expectedValue)
	}
}
