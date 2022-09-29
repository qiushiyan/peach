package eval

import (
	"fmt"
	"reflect"

	"github.com/qiushiyan/qlang/pkg/ast"
	"github.com/qiushiyan/qlang/pkg/object"
)

func evalVectorLiteral(node ast.Node, env *object.Env) object.Object {
	vectorNode := node.(*ast.VectorLiteral)
	elements := evalExpressions(vectorNode.Elements, env)

	firstElement := elements[0]
	if len(elements) == 1 && isError(firstElement) {
		return elements[0]
	}

	switch firstElement.(type) {
	case *object.Number:
		if sameType(elements, reflect.TypeOf(firstElement)) {
			fmt.Println(reflect.TypeOf(firstElement))
			return &object.NumericVector{Elements: elements}
		}
	case *object.String:
		if sameType(elements, reflect.TypeOf(firstElement)) {
			return &object.CharacterVector{Elements: elements}
		}
	case *object.Boolean:
		if sameType(elements, reflect.TypeOf(firstElement)) {
			return &object.LogicalVector{Elements: elements}
		}
	}
	return &object.Vector{Elements: elements}
}

func sameType(values []object.Object, t reflect.Type) bool {
	if len(values) == 0 {
		return true
	}
	for _, v := range values[1:] {
		if reflect.TypeOf(v) != t {
			return false
		}
	}
	return true
}
