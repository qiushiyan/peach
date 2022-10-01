package eval

import "testing"

func TestCall(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{"let identity = fn(x) { x }; identity(5);", 5},
		{"let identity = fn(x) { return x }; identity(5);", 5},
		{"let double = fn(x) { x * 2; }; double(5);", 10},
		{"let add = fn(x, y = 1) { x + y }; add(5, 5);", 10},
		{"let add = fn(x, y = 1) { x + y }; add(5 + 5);", 11},
		// {"let add = fn(x, y = 1, z = 2) { x + y + 2 * z }; add(5, z = 1);", 9},
		{"fn(x) { x }(5)", 5},
	}

	for _, tt := range tests {
		testNumberObject(t, testEval(tt.input), tt.expected)
	}
}
