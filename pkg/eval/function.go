package eval

import (
	"github.com/qiushiyan/qlang/pkg/ast"
	"github.com/qiushiyan/qlang/pkg/object"
)

func evalFunctionLiteral(node *ast.FunctionLiteral, env *object.Env) object.Object {
	return &object.Function{
		Parameters:            node.Parameters,
		RequiredParametersNum: len(node.Parameters),
		Body:                  node.Body,
		Env:                   env,
	}
}
