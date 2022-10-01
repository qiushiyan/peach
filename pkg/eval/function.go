package eval

import (
	"github.com/qiushiyan/qlang/pkg/ast"
	"github.com/qiushiyan/qlang/pkg/object"
)

func evalFunctionLiteral(node *ast.FunctionLiteral, env *object.Env) object.Object {
	return &object.Function{
		Parameters:            node.Parameters,
		Defaults:              node.Defaults,
		RequiredParametersNum: len(node.Parameters) - len(node.Defaults),
		Body:                  node.Body,
		Env:                   env,
	}
}
