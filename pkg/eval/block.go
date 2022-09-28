package eval

import (
	"github.com/qiushiyan/qlang/pkg/ast"
	"github.com/qiushiyan/qlang/pkg/object"
)

func evalBlockStatement(node *ast.BlockStatement) object.Object {
	var result object.Object
	for _, statement := range node.Statements {
		result = Eval(statement)
		if isError(result) {
			return result
		}
		if result != nil && result.Type() == object.RETURN_OBJ {
			return result
		}
	}
	return result
}
