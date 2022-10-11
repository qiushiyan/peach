package eval

import "testing"

func TestIdentifierNesetedEnv(t *testing.T) {
	input := `
	a <- 1
	if (true) {
		for (x in [1]) {
			if (true) {
				a
			}
		}
	}
	`
	evaluated := testEval(input)
	testNumberObject(t, evaluated, 1)
}
