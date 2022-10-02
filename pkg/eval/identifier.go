package eval

import (
	"github.com/qiushiyan/qlang/pkg/ast"
	"github.com/qiushiyan/qlang/pkg/object"
	"github.com/qiushiyan/qlang/pkg/std"
)

func evalIdentifier(node *ast.Identifier, env *object.Env) object.Object {
	if val, ok := env.Get(node.Value); ok {
		return val
	}
	if builtin, ok := std.Builtins[node.Value]; ok {
		return builtin
	}
	return object.NewError("identifier not found: " + node.Value)
}
