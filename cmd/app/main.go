package main

import (
	"fmt"
	"os"
	"strconv"
	"task-cli/internal/handlers"
)

type actionHandler func(args []string) error

func main() {
	commandLineArgs := os.Args

	if len(commandLineArgs) < 2 {
		fmt.Println("action and arguments param required.")
		fmt.Println("example: task-cli add 1")
		os.Exit(1)
	}

	availableActions := map[string]actionHandler{
		"add":              handlers.AddHandler,
		"update":           handlers.UpdateHandler,
		"delete":           handlers.DeleteHandler,
		"mark-in-progress": handlers.MarkInProgressHandler,
		"mark-done":        handlers.MarkDoneHandler,
		"list":             handlers.ListHandler,
	}

	actionName := commandLineArgs[1]

	actionArgs := commandLineArgs[2:]

	handler, exists := availableActions[actionName]
	if !exists {
		fmt.Printf("unsupported action: %v\n", actionName)
		os.Exit(1)
	}

	if err := handler(actionArgs); err != nil {
		fmt.Printf("error: %v\n\n", err)
		os.Exit(1)
	}
}

func validateID(id string) {
	_, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("task id must be an integer..")
		return
	}
}
