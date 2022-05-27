package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args
	if len(args) == 1 {
		fmt.Println("Please provide an argument")
		return
	}

	handleUserInput(os.Args[1:])
}
