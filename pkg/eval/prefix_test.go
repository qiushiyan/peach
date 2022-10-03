package eval

import "testing"

func TestBangOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!true", false},
		{"!false", true},
		{"!5", false},
		{"!!true", true},
		{"!!false", false},
		{"!!5", true},
		{"!null", true},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		if tt.input != "!null" {
			testBooleanObject(t, evaluated, tt.expected)
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func TestMinusPrefixOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{"-5", -5},
		{"-10", -10},
		{"--5", 5},
		{"-(-5)", 5},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testNumberObject(t, evaluated, tt.expected)
	}
}

func TestPlusPrefixOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{"+5", 5},
		{"+10", 10},
		{"++5", 5},
		{"+(-5)", -5},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testNumberObject(t, evaluated, tt.expected)
	}
}
