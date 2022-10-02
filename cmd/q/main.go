package main

import (
	"fmt"
	"io"
	"os"
	"os/user"

	"github.com/qiushiyan/qlang/pkg/object"
	"github.com/qiushiyan/qlang/pkg/repl"
	"github.com/qiushiyan/qlang/pkg/std"
)

func main() {
	std.RegisterStd()
	if len(os.Args) < 2 {
		// repl
		user, err := user.Current()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Hello %s! Welcome to the Q language!\n",
			user.Username)
		repl.Start(os.Stdin, os.Stdout)
		return
	}

	// the scripting interface
	file := os.Args[1]

	buf, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	evaluated := repl.Evaluate(os.Stdout, string(buf), object.NewEnv())
	if evaluated != nil {
		if evaluated.Type() != object.NULL_OBJ {
			io.WriteString(os.Stdout, evaluated.Inspect())
			io.WriteString(os.Stdout, "\n")
		}
	}

}
