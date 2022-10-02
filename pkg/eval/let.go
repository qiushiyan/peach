package eval

import (
	"github.com/qiushiyan/qlang/pkg/ast"
	"github.com/qiushiyan/qlang/pkg/object"
)

func evalLetStatement(node *ast.LetStatement, env *object.Env) object.Object {
	val := Eval(node.Value, env)
	if object.IsError(val) {
		return val
	}
	env.Set(node.Name.Value, val)
	return NULL
}
