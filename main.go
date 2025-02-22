package main

import (
	"bolt/repl"
	"fmt"
	"os"
	"os/user"
)

// Start the Bolt REPL
//   - Get the current user
//   - Print a welcome message
//   - Start the REPL
func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! Welcome to Bolt ⚡️\n", user.Username)
	fmt.Printf("Type a command and press Enter to execute it.\n")
	repl.Start(os.Stdin, os.Stdout)
}
