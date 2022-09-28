package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/qiushiyan/qlang/pkg/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! Welcome to the Q language!\n",
		user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
