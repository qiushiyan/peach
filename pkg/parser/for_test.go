package parser

import (
	"strings"
	"testing"

	"github.com/qiushiyan/qlang/pkg/ast"
	"github.com/qiushiyan/qlang/pkg/lexer"
)

func TestForExpression(t *testing.T) {
	input := `for (x in y) { x }`

	l := lexer.New(strings.NewReader(input))
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain %d statements. got=%d\n",
			1, len(program.Statements))
	}

	statement, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
			program.Statements[0])
	}

	exp, ok := statement.Expression.(*ast.ForExpression)
	if !ok {
		t.Fatalf("statement.Expression is not ast.ForExpression. got=%T", statement.Expression)
	}

	if !testIdentifier(t, exp.Iterator, "x") {
		return
	}

	if !testIdentifier(t, exp.IterValues, "y") {
		return
	}

	forStatements := exp.Statements.Statements
	if len(forStatements) != 1 {
		t.Errorf("for does not contain 1 statement. got=%d\n",
			len(forStatements))
	}

	forStatement, ok := forStatements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("for statements[0] is not ast.ExpressionStatement. got=%T",
			forStatement)
	}

	if !testIdentifier(t, forStatement.Expression, "x") {
		return
	}
}
