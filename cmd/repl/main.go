package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/qiushiyan/peach/pkg/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! Welcome to the Peach 🍑 programming language!\n",
		user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
