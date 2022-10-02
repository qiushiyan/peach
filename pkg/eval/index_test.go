package eval

import "testing"

func TestIndexExpressopm(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"[1, 2, 3][1]", "1"},
		{"[1, 2, 3][3]", "3"},
		{"[1, 2, 3][1:]", "[1, 2, 3]"},
		{"[1, 2, 3][1:3]", "[1, 2, 3]"},
		{"[1, 2, 3][:3]", "[1, 2, 3]"},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		if evaluated.Inspect() != tt.expected {
			t.Errorf("expected %q, got %q", tt.expected, evaluated.Inspect())
		}
	}
}
