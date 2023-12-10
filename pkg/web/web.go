package web

import (
	"strings"

	"github.com/qiushiyan/qlang/pkg/eval"
	"github.com/qiushiyan/qlang/pkg/lexer"
	"github.com/qiushiyan/qlang/pkg/object"
	"github.com/qiushiyan/qlang/pkg/parser"
)

type errorString struct {
	s string
}

func (e errorString) Error() string {
	return e.s
}

func Evaluate(input string, env *object.Env) (object.Object, error) {
	l := lexer.New(strings.NewReader(input))
	p := parser.New(l)

	program := p.ParseProgram()
	// parsing errors
	if len(p.Errors()) != 0 {
		allError := strings.Join(p.Errors(), "\n")
		return nil, errorString{allError}
	}

	obj := eval.Eval(program, env)
	// evaluation error
	if obj.Type() == object.ERROR_OBJ {
		return nil, errorString{obj.Inspect()}
	} else {
		return obj, nil
	}
}
