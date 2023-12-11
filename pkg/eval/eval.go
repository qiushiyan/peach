package eval

import (
	"strings"

	"github.com/qiushiyan/qlang/pkg/ast"
	"github.com/qiushiyan/qlang/pkg/lexer"
	"github.com/qiushiyan/qlang/pkg/object"
	"github.com/qiushiyan/qlang/pkg/parser"
)

func Eval(node ast.Node, env *object.Env) object.Object {

	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node, env)
	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)
	case *ast.Identifier:
		return evalIdentifier(node, env)
	case *ast.NumberLiteral:
		return evalNumberLiteral(node.Value)
	case *ast.StringLiteral:
		return evalStringLiteral(node.Value)
	case *ast.Boolean:
		return evalBoolean(node.Value)
	case *ast.ForExpression:
		return evalForExpression(node, env)
	case *ast.RangeExpression:
		return evalRangeExpression(node, env)
	case *ast.VectorLiteral:
		return evalVectorLiteral(node, env)
	case *ast.DictLiteral:
		return evalDictLiteral(node, env)
	case *ast.IndexExpression:
		return evalIndexExpression(node, env)
	case *ast.FunctionLiteral:
		return evalFunctionLiteral(node, env)
	case *ast.CallExpression:
		return evalCallExpression(node, env)
	case *ast.Null:
		return evalNull()
	case *ast.ReturnStatement:
		return evalReturnStatement(node, env)
	case *ast.LetStatement:
		return evalLetStatement(node, env)
	case *ast.AssignExpression:
		return evalAssignExpression(node, env)
	case *ast.PrefixExpression:
		return evalPrefixExpression(node, env)
	case *ast.InfixExpression:
		if node.Operator == "|>" {
			call := evalPipeExpression(node, env)
			if call == nil {
				return &object.Error{Message: "right hide side of |> should be a function call or function"}
			}
			return evalCallExpression(call, env)
		} else {
			left := Eval(node.Left, env)
			if object.IsError(left) {
				return left
			}
			right := Eval(node.Right, env)
			if object.IsError(right) {
				return right
			}
			return evalInfixExpression(node.Operator, left, right)
		}
	case *ast.IfExpression:
		return evalIfExpression(node, env)
	case *ast.BlockStatement:
		return evalBlockStatement(node, env)
	}

	return nil
}

func evalProgram(program *ast.Program, env *object.Env) object.Object {
	if len(program.Statements) == 0 {
		// when the program is empty, e.g., a file with only comments
		return &object.Null{}
	}

	var result object.Object

	for _, statement := range program.Statements {
		result = Eval(statement, env)
		if object.IsError(result) {
			return result
		}
		if returnValue, ok := result.(*object.ReturnValue); ok {
			return returnValue.Value
		}
	}

	return result
}

func evalExpressions(exps []ast.Expression, env *object.Env) []object.Object {
	var result []object.Object
	for _, e := range exps {
		evaluated := Eval(e, env)
		if object.IsError(evaluated) {
			return []object.Object{evaluated}
		}
		result = append(result, evaluated)
	}
	return result
}

func testEval(input string) object.Object {
	l := lexer.New(strings.NewReader(input))
	p := parser.New(l)
	env := object.NewEnv()
	program := p.ParseProgram()
	// no registeration of stdlib in testEval
	return Eval(program, env)
}
