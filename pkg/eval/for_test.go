package eval

import (
	"testing"

	"github.com/qiushiyan/qlang/pkg/object"
)

func TestForExpression(t *testing.T) {
	input := `
	x = vector(3)
	for (i in [1, 2, 3]) {
		x[i] = i
	}
	x
	`
	evaluated := testEval(input)
	result, ok := evaluated.(*object.Vector)
	if !ok {
		t.Fatalf("object is not Vector. got=%T (%+v)", evaluated, evaluated)
	}
	if len(result.Elements) != 3 {
		t.Fatalf("vector has wrong num of elements. got=%d",
			len(result.Elements))
	}
	testNumberObject(t, result.Elements[0], 1)
	testNumberObject(t, result.Elements[1], 2)
	testNumberObject(t, result.Elements[2], 3)
}
