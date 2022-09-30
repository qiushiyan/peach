package eval

import (
	"github.com/qiushiyan/qlang/pkg/ast"
	"github.com/qiushiyan/qlang/pkg/object"
)

func evalReturnStatement(node *ast.ReturnStatement, env *object.Env) object.Object {
	val := Eval(node.Value, env)
	if object.IsError(val) {
		return val
	}
	return &object.ReturnValue{Value: val}
}
