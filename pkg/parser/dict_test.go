package parser

import (
	"strings"
	"testing"

	"github.com/qiushiyan/qlang/pkg/ast"
	"github.com/qiushiyan/qlang/pkg/lexer"
)

func TestParsingDictLiteralsWithExpressions(t *testing.T) {
	input := `{"one": 0 + 1, "two": 10 - 8, "three": 15 / 5, "four": 1:10}`
	l := lexer.New(strings.NewReader(input))
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	statement := program.Statements[0].(*ast.ExpressionStatement)
	DictLiteral, ok := statement.Expression.(*ast.DictLiteral)
	if !ok {
		t.Fatalf("exp is not ast.DictLiteral. got=%T", statement.Expression)
	}
	if len(DictLiteral.Pairs) != 4 {
		t.Errorf("Dict.Pairs has wrong length. got=%d", len(DictLiteral.Pairs))
	}

	tests := map[string]func(ast.Expression){
		"one": func(e ast.Expression) {
			testInfixExpression(t, e, 0, "+", 1)
		},
		"two": func(e ast.Expression) {
			testInfixExpression(t, e, 10, "-", 8)
		},
		"three": func(e ast.Expression) {
			testInfixExpression(t, e, 15, "/", 5)
		},
		"four": func(e ast.Expression) {
			testRangeExpression(t, e, 1, 10)
		},
	}
	for key, value := range DictLiteral.Pairs {
		literal, ok := key.(*ast.StringLiteral)
		if !ok {
			t.Errorf("key is not ast.StringLiteral. got=%T", key)
			continue
		}
		testFunc, ok := tests[literal.String()]
		if !ok {
			t.Errorf("No test function for key %q found", literal.String())
			continue
		}
		testFunc(value)
	}

}
