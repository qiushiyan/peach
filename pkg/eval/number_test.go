package eval

import (
	"testing"

	"github.com/qiushiyan/peach/pkg/object"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{"5", 5},
		{"10", 10}}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testNumberObject(t, evaluated, tt.expected)
	}
}

func testNumberObject(t *testing.T, obj object.Object, expected float64) bool {
	result, ok := obj.(*object.Number)
	if !ok {
		t.Errorf("object is not number. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%v, want=%v",
			result.Value, expected)
		return false
	}
	return true
}
