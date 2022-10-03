package std

import (
	"fmt"
	"math/rand"

	"github.com/qiushiyan/qlang/pkg/object"
)

func random(env *object.Env, args ...object.Object) object.Object {
	if len(args) == 0 {
		return newNumericVector([]object.Object{&object.Number{Value: rand.Float64()}})
	}
	num, ok := args[0].(*object.Number)
	if !ok {
		return object.NewError("invalid first argument in function `random()`, must be a number, got %s\nUsage: random(n, low, high)", args[0].Type())
	}
	var els = make([]object.Object, 0, int(num.Value))
	var low, high float64 = 0, 1
	if len(args) >= 2 {
		if args[1].Type() != object.NUMBER_OBJ {
			fmt.Println("Usage: random(n, low, high)")
			return object.NewError("invalid argument in function `random()`, must be a number, got %s\nUsage: random(n, low, high)", args[0].Type())
		}
		low = args[1].(*object.Number).Value
	}
	if len(args) >= 3 {
		if args[2].Type() != object.NUMBER_OBJ {
			fmt.Println("Usage: random(n, low, high)")
			return object.NewError("invalid argument in function `random()`, must be a number, got %s\nUsage: random(n, low, high)", args[0].Type())
		}
		high = args[2].(*object.Number).Value
	}
	for i := 0; i < int(num.Value); i++ {
		els = append(els, &object.Number{Value: rand.Float64()*(high-low) + low})
	}
	return newNumericVector(els)
}

func newNumericVector(els []object.Object) object.Object {
	return &object.NumericVector{BaseVector: object.BaseVector{Elements: els}}
}
