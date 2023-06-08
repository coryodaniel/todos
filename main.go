package main

import (
	"fmt"
	"os"

	"github.com/coryodaniel/todo/cmd"
)

func main() {
	if err := cmd.Execute(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
