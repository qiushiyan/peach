package eval

import (
	"fmt"

	"github.com/qiushiyan/qlang/pkg/ast"
	"github.com/qiushiyan/qlang/pkg/object"
)

func evalBlockStatement(node *ast.BlockStatement, env *object.Env) object.Object {
	var result object.Object
	for _, stmt := range node.Statements {
		fmt.Println(stmt.String())
	}
	for _, statement := range node.Statements {
		result = Eval(statement, env)
		if object.IsError(result) {
			return result
		}
		if result != nil && result.Type() == object.RETURN_OBJ {
			return result
		}
	}
	return result
}
