package parser

import (
	"strings"
	"testing"

	"github.com/qiushiyan/peach/pkg/ast"
	"github.com/qiushiyan/peach/pkg/lexer"
)

func TestBlockStatementParsing(t *testing.T) {
	input := `{ x + 1; y + 2; z + 3; }`

	l := lexer.New(strings.NewReader(input))
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statements. got=%d", len(program.Statements))
	}

	blockStatement, ok := program.Statements[0].(*ast.BlockStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.BlockStatement. got=%T", program.Statements[0])
	}

	if len(blockStatement.Statements) != 3 {
		t.Fatalf("blockStatement.Statements does not contain %d statements. got=%d\n",
			2, len(blockStatement.Statements))
	}

	expectedExprs := []struct {
		left     interface{}
		operator string
		right    interface{}
	}{
		{"x", "+", 1.0},
		{"y", "+", 2.0},
		{"z", "+", 3.0},
	}
	for i, statement := range blockStatement.Statements {
		if _, ok := statement.(*ast.ExpressionStatement); !ok {
			t.Fatalf("statement is not ast.ExpressionStatement. got=%T",
				statement)
		}

		expectedExpr := expectedExprs[i]
		if expr := statement.(*ast.ExpressionStatement).Expression; !testInfixExpression(t, expr, expectedExpr.left, expectedExpr.operator, expectedExpr.right) {
			t.Fatalf("statement is not ast.ExpressionStatement. got=%T",
				statement)
		}
	}
	// statements := blockStatement.Statements

}
