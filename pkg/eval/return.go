package eval

import (
	"github.com/qiushiyan/qlang/pkg/ast"
	"github.com/qiushiyan/qlang/pkg/object"
)

func evalReturnStatement(node *ast.ReturnStatement) object.Object {
	val := Eval(node.Value)
	if isError(val) {
		return val
	}
	return &object.ReturnValue{Value: val}
}
