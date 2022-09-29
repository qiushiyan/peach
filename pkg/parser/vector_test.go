package parser

import (
	"strings"
	"testing"

	"github.com/qiushiyan/qlang/pkg/ast"
	"github.com/qiushiyan/qlang/pkg/lexer"
)

func TestParsingVectorLiterals(t *testing.T) {
	input := "[1, 2 * 2, 3 + 3]"
	l := lexer.New(strings.NewReader(input))
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	statement, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}
	Vector, ok2 := statement.Expression.(*ast.VectorLiteral)
	if !ok2 {
		t.Fatalf("exp not ast.VectorLiteral. got=%T", statement.Expression)
	}

	if len(Vector.Elements) != 3 {
		t.Fatalf("len(Vector.Elements) not 3. got=%d", len(Vector.Elements))
	}
	testNumberLiteral(t, Vector.Elements[0], 1)
	testInfixExpression(t, Vector.Elements[1], 2, "*", 2)
	testInfixExpression(t, Vector.Elements[2], 3, "+", 3)
}
