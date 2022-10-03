package eval

import "testing"

func TestEnclosedEnvironment(t *testing.T) {
	input := `
	let first = 10;
	let second = 10;
	let third = 10;

	let ourFunction = fn(first) {
		let second = 20

		first + second + third
	};

	ourFunction(20) + first + second
	`
	result := testEval(input)
	testNumberObject(t, result, 70)
}

func TestClosures(t *testing.T) {
	input := `
	let newAdder = fn(x) {
		fn(y) { x + y };
	};

	let addTwo = newAdder(2);
	addTwo(2)
	`

	testNumberObject(t, testEval(input), 4)
}

func TestDefaultArguments(t *testing.T) {
	input := `
	let foo = fn(a, b=1) { a + b };
	foo(1)
	`

	testNumberObject(t, testEval(input), 2)
}

func TestNamedArguments(t *testing.T) {
	input := `
	let foo = fn(a, b = 1, z = 2) { a + b + z * 2 };
	foo(1, z = 3)
	`

	testNumberObject(t, testEval(input), 8)
}
