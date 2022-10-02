package eval

import (
	"testing"

	"github.com/qiushiyan/qlang/pkg/object"
)

func TestDictLiterals(t *testing.T) {
	input := `let two = "two"
	{
		"one": 10 - 9,
		two: 1 + 1,
		"thr" + "ee": 6 / 2,
		4: 4,
		true: 5,
		false: 6
	}`

	evaluated := testEval(input)
	result, ok := evaluated.(*object.Dict)
	if !ok {
		t.Fatalf("Eval didn't return Dict. got=%T (%+v)", evaluated, evaluated)
	}

	expected := map[object.DictKey]float64{
		(&object.String{Value: "one"}).Hash():   1,
		(&object.String{Value: "two"}).Hash():   2,
		(&object.String{Value: "three"}).Hash(): 3,
		(&object.Integer{Value: 4}).Hash():      4,
		object.TRUE.Hash():                      5,
		object.FALSE.Hash():                     6,
	}

	if len(result.Pairs) != len(expected) {
		t.Fatalf("Dict has wrong num of pairs. got=%d", len(result.Pairs))
	}

	for expectedKey, expectedValue := range expected {
		pair, ok := result.Pairs[expectedKey]
		if !ok {
			t.Errorf("no pair for given key in Pairs")
		}

		testNumberObject(t, pair.Value, expectedValue)
	}
}

func TestDictIndexExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			`{"foo": 5}["foo"]`,
			5,
		},
		{
			`{foo: 5}["foo"]`,
			5,
		},
		{
			`{"foo": 5}["bar"]`,
			nil,
		},
		{
			`let key = "foo"; {"foo": 5}[key]`,
			5,
		},
		{
			`{}["foo"]`,
			nil,
		},
		{
			`{5: 5}[5]`,
			5,
		},
		{
			`{true: 5}[true]`,
			5,
		},
		{
			`{false: 5}[false]`,
			5,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		num, ok := tt.expected.(int)
		if ok {
			testNumberObject(t, evaluated, float64(num))
		} else {
			testNullObject(t, evaluated)
		}
	}
}
