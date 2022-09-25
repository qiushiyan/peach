package repl

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/qiushiyan/peach/pkg/lexer"
	"github.com/qiushiyan/peach/pkg/parser"
)

const PROMPT = "\U0000279C  "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(strings.NewReader(line))
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}

}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "An error occurred!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
