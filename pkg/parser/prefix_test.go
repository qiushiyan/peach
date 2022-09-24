package parser

import (
	"fmt"
	"strings"
	"testing"

	"github.com/qiushiyan/peach/pkg/ast"
	"github.com/qiushiyan/peach/pkg/lexer"
)

func TestParsingPrefixExpressions(t *testing.T) {
	prefixTests := []struct {
		input    string
		operator string
		number   float64
	}{
		{"!5;", "!", 5},
		{"-15.5;", "-", 15.5},
	}

	for _, tt := range prefixTests {
		l := lexer.New(strings.NewReader(tt.input))
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
		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("stmt is not ast.PrefixExpression. got=%T", stmt.Expression)
		}
		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not '%s'. got=%s",
				tt.operator, exp.Operator)
			if !testNumberLiteral(t, exp.Right, tt.number) {
				return
			}
		}
	}
}

func testNumberLiteral(t *testing.T, nl ast.Expression, value float64) bool {
	number, ok := nl.(*ast.NumberLiteral)
	if !ok {
		t.Errorf("il not *ast.IntegerLiteral. got=%T", nl)
		return false
	}
	if number.Value != value {
		t.Errorf("number.Value not %f. got=%f", value, number.Value)
		return false
	}
	if number.TokenLiteral() != fmt.Sprintf("%f", value) {
		t.Errorf("integ.TokenLiteral not %f. got=%s", value,
			number.TokenLiteral())
		return false
	}
	return true
}
