package repl

import (
	"bufio"
	"fmt"
	"go_interp/lexer"
	"go_interp/token"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprint(out, "\n", tok)
		}

    fmt.Fprint(out, "\n\n")
	}
}
