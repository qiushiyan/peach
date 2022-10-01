package eval

import (
	"testing"

	"github.com/qiushiyan/qlang/pkg/object"
)

func TestErrorHandling(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{

		{
			"5 + true",
			"operaor + needs values of the same type, got NUMBER + BOOLEAN"},
		{
			"\"string\" + 1",
			"operaor + needs values of the same type, got STRING + NUMBER",
		},
		{
			"\"string\" - \"string\"",
			"invalid operation STRING - STRING",
		},
		{
			"-true",
			"invalid operator - for type BOOLEAN",
		},
		{
			"if (10 > 1) { true + false; }",
			"invalid operation BOOLEAN + BOOLEAN",
		},
		{`if (10 > 1) {
		  	if (10 > 1) {
				return true + false;
			}
			return 1
		}`, "invalid operation BOOLEAN + BOOLEAN",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Errorf("no error object returned. got=%T(%+v)",
				evaluated, evaluated)
			continue
		}
		if errObj.Message != tt.expectedMessage {
			t.Errorf("wrong error message. expected=%q, got=%q",
				tt.expectedMessage, errObj.Message)
		}
	}
}
