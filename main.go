package main

import (
	"fmt"
	"os"

	"github.com/thr3m/nojs/cli"
)

func main() {

	args := os.Args
	if len(args) == 1 {
		fmt.Println("Please provide an argument")
		return
	}

	cli.HandleUserInput(os.Args[1:])
}
