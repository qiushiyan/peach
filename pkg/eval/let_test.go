package eval

import (
	"testing"

	"github.com/qiushiyan/qlang/pkg/object"
)

func TestEvalLetStatement(t *testing.T) {
	input := `let a = 5; a`

	evaluated := testEval(input)
	testNumberObject(t, evaluated, 5)
}

func TestEvalLetStatementOuter(t *testing.T) {
	input := `if (true) { let a = 5}; a`

	evaluated := testEval(input)
	err, ok := evaluated.(*object.Error)
	if !ok {
		t.Errorf("object is not Error. got=%T (%+v)", evaluated, evaluated)
	}
	if err.Message != "identifier not found: a" {
		t.Errorf("wrong error message. got=%q", err.Message)
	}
}
