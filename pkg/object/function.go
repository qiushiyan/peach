package object

import (
	"bytes"
	"strings"

	"github.com/qiushiyan/qlang/pkg/ast"
)

type Function struct {
	Parameters            []*ast.Identifier
	Defaults              map[string]ast.Expression
	RequiredParametersNum int
	Body                  *ast.BlockStatement
	Env                   *Env
}

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
	var out bytes.Buffer
	params := make([]string, len(f.Parameters))
	for i := range f.Parameters {
		params[i] = f.Parameters[i].String()
	}
	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n\t")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")
	return out.String()
}
