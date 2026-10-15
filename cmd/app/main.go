package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		os.Exit(1)
	}

	actionName := args[1]

	switch actionName {
	case "add":
		if len(args) < 3 {
			fmt.Println("add command requires a task name..")
			return
		}
		// add
	case "update":
		if len(args) < 4 {
			fmt.Println("update command requires an index and a task name..")
			return
		}
		// update
	}

}
