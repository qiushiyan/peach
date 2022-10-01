package parser

import (
	"strings"
	"testing"

	"github.com/qiushiyan/qlang/pkg/ast"
	"github.com/qiushiyan/qlang/pkg/lexer"
)

func TestFunctionParameterParsing(t *testing.T) {
	tests := []struct {
		input          string
		expectedParams []string
	}{
		{input: "fn() {};", expectedParams: []string{}},
		{input: "fn(x) {};", expectedParams: []string{"x"}},
		{input: "fn(x, y, z) {};", expectedParams: []string{"x", "y", "z"}},
	}
	for _, tt := range tests {
		l := lexer.New(strings.NewReader(tt.input))
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)
		statement := program.Statements[0].(*ast.ExpressionStatement)
		function := statement.Expression.(*ast.FunctionLiteral)
		if len(function.Parameters) != len(tt.expectedParams) {
			t.Errorf("length parameters wrong. want %d, got=%d\n",
				len(tt.expectedParams), len(function.Parameters))
		}
		for i, identifier := range tt.expectedParams {
			testIdentifier(t, function.Parameters[i], identifier)
		}
	}
}

func TestFunctionDefaultParameterParsing(t *testing.T) {
	input := `fn(x, y = 1, z = 2) {}`
	l := lexer.New(strings.NewReader(input))
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	statement := program.Statements[0].(*ast.ExpressionStatement)
	function := statement.Expression.(*ast.FunctionLiteral)

	if len(function.Parameters) != 3 {
		t.Errorf("number of parameters wrong. want=%d, got=%d\n",
			3, len(function.Parameters))
	}

	if !testIdentifier(t, function.Parameters[0], "x") {
		return
	}

	if !testIdentifier(t, function.Parameters[1], "y") {
		return
	}
	if !testIdentifier(t, function.Parameters[2], "z") {
		return
	}

	if len(function.Defaults) != 2 {
		t.Errorf("number of default values wrong. want=%d, got=%d\n",
			2, len(function.Defaults))
	}

	if !testLiteralExpression(t, function.Defaults["y"], 1) {
		t.Errorf("default value for y wrong. want=%v, got=%v\n",
			1, function.Defaults["y"].(*ast.NumberLiteral).Value)
	}

	if !testLiteralExpression(t, function.Defaults["z"], 2) {
		t.Errorf("default value for z wrong. want=%v, got=%v\n",
			2, function.Defaults["z"].(*ast.NumberLiteral).Value)
	}
}

func TestFunctionLiteralParsing(t *testing.T) {
	tests := []string{
		"fn(x, y) { x + y }",
		"fn(x, y) x + y",
	}

	for _, input := range tests {
		l := lexer.New(strings.NewReader(input))
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statements. got=%d\n",
				1, len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
				program.Statements[0])
		}

		function, ok := stmt.Expression.(*ast.FunctionLiteral)
		if !ok {
			t.Fatalf("stmt.Expression is not ast.FunctionLiteral. got=%T",
				stmt.Expression)
		}

		if len(function.Parameters) != 2 {
			t.Fatalf("function literal parameters wrong. want 2, got=%d\n",
				len(function.Parameters))
		}

		testLiteralExpression(t, function.Parameters[0], "x")
		testLiteralExpression(t, function.Parameters[1], "y")

		if len(function.Body.Statements) != 1 {
			t.Fatalf("function.Body.Statements has not 1 statements. got=%d\n",
				len(function.Body.Statements))
		}

		returnStatement, ok := function.Body.Statements[0].(*ast.ReturnStatement)
		if !ok {
			t.Fatalf("function body stmt is not ast.ReturnStatement. got=%T",
				function.Body.Statements[0])
		}

		testInfixExpression(t, returnStatement.Value, "x", "+", "y")
	}

}
