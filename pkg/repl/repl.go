package repl

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/qiushiyan/peach/pkg/lexer"
	"github.com/qiushiyan/peach/pkg/token"
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

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}

}
