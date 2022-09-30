package eval

import (
	"testing"

	"github.com/qiushiyan/qlang/pkg/object"
)

func TestArrayLiterals(t *testing.T) {
	input := "[1, 2 * 2, 3 + 3]"
	evaluated := testEval(input)
	result, ok := evaluated.(*object.NumericVector)
	if !ok {
		t.Fatalf("object is not NumericVector. got=%T (%+v)", evaluated, evaluated)
	}
	if len(result.Elements) != 3 {
		t.Fatalf("vector has wrong num of elements. got=%d",
			len(result.Elements))
	}
	testNumberObject(t, result.Elements[0], 1)
	testNumberObject(t, result.Elements[1], 4)
	testNumberObject(t, result.Elements[2], 6)
}
