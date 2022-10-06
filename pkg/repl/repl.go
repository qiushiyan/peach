package repl

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/qiushiyan/qlang/pkg/eval"
	"github.com/qiushiyan/qlang/pkg/lexer"
	"github.com/qiushiyan/qlang/pkg/object"
	"github.com/qiushiyan/qlang/pkg/parser"
)

const (
	PROMPT = "\U0000279C  "
)

type Config struct {
	interactive bool
}

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnv()
	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		evaluated := Evaluate(out, line, env)
		if evaluated != nil {
			if evaluated.Type() != object.NULL_OBJ {
				io.WriteString(out, evaluated.Inspect())
				io.WriteString(out, "\n")
			}
		}
	}

}

func Evaluate(out io.Writer, input string, env *object.Env) object.Object {
	l := lexer.New(strings.NewReader(input))
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		printParserErrors(out, p.Errors())
		return nil
	}

	return eval.Eval(program, env)
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "An error occurred!\n")
	io.WriteString(out, " syntax error:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
