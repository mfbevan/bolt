package repl

import (
	"bolt/lexer"
	"bolt/token"
	"bufio"
	"fmt"
	"io"
)

// The REPL prompt is prepended to each input line and is used to indicate that the REPL is ready to accept input
const PROMPT = "⚡️> "

// Start the Bolt REPL
//   - Read input from the user
//   - Tokenize the input
//   - Print the tokens
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
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
