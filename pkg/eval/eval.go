package eval

import (
	"strings"

	"github.com/qiushiyan/peach/pkg/ast"
	"github.com/qiushiyan/peach/pkg/lexer"
	"github.com/qiushiyan/peach/pkg/object"
	"github.com/qiushiyan/peach/pkg/parser"
)

var (
	NULL  = &object.Null{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.NumberLiteral:
		return evalNumberLiteral(node.Value)
	case *ast.StringLiteral:
		return &object.String{Value: node.Value}
	case *ast.Boolean:
		return evalBoolean(node.Value)
	case *ast.Null:
		return evalNull()
	case *ast.PrefixExpression:
		right := Eval(node.Right)
		return evalPrefixExpression(node.Operator, right)
	case *ast.InfixExpression:
		left := Eval(node.Left)
		right := Eval(node.Right)
		return evalInfixExpression(node.Operator, left, right)
	}

	return nil
}

func evalStatements(statements []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range statements {
		result = Eval(statement)
	}

	return result
}

func testEval(input string) object.Object {
	l := lexer.New(strings.NewReader(input))
	p := parser.New(l)
	program := p.ParseProgram()
	return Eval(program)
}

func evalNumberLiteral(value float64) object.Object {
	return &object.Number{Value: value}
}

func evalNull() *object.Null {
	return NULL
}

func evalBoolean(value bool) *object.Boolean {
	if value {
		return TRUE
	} else {
		return FALSE
	}
}

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusPrefixOperatorExpression(right)
	default:
		return NULL
	}
}

func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	// !null is still null
	case NULL:
		return NULL
	default:
		// !nonzero number = false
		// !zero number = true
		if number, ok := right.(*object.Number); ok {
			if number.Value != 0 {
				return FALSE
			} else {
				return TRUE
			}
		}
		return FALSE
	}
}

func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	if right.Type() != object.NUMBER_OBJ {
		return NULL
	}
	value := right.(*object.Number).Value
	return &object.Number{Value: -value}
}

func evalInfixExpression(operator string, left object.Object, right object.Object) object.Object {
	switch {
	case left.Type() == object.NUMBER_OBJ && right.Type() == object.NUMBER_OBJ:
		return evalNumberInfixExpression(operator, left, right)
	case left.Type() == object.STRING_OBJ && right.Type() == object.STRING_OBJ:
		return evalStringInfixExpression(operator, left, right)
	case operator == "==":
		return evalBoolean(left == right)
	case operator == "!=":
		return evalBoolean(left != right)
	case left.Type() != right.Type():
		return NULL
	default:
		return NULL
	}
}

func evalNumberInfixExpression(operator string, left object.Object, right object.Object) object.Object {
	leftVal := left.(*object.Number).Value
	rightVal := right.(*object.Number).Value

	switch operator {
	case "+":
		return &object.Number{Value: leftVal + rightVal}
	case "-":
		return &object.Number{Value: leftVal - rightVal}
	case "*":
		return &object.Number{Value: leftVal * rightVal}
	case "/":
		return &object.Number{Value: leftVal / rightVal}
	case "<":
		return evalBoolean(leftVal < rightVal)
	case ">":
		return evalBoolean(leftVal > rightVal)
	case "==":
		return evalBoolean(leftVal == rightVal)
	case "!=":
		return evalBoolean(leftVal != rightVal)
	default:
		return NULL
	}
}

func evalStringInfixExpression(operator string, left object.Object, right object.Object) object.Object {
	leftVal := left.(*object.String).Value
	rightVal := right.(*object.String).Value

	switch operator {
	case "+":
		return &object.String{Value: leftVal + rightVal}
	default:
		return NULL
	}
}
