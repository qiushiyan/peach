package eval

import "testing"

func TestPipeOpeartor(t *testing.T) {
	input := `
	add = fn(x, y, z) { x + y + z};
	1 |> add(2, 3)
	`

	testNumberObject(t, testEval(input), 6)
}
