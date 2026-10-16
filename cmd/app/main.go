package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("action and arguments param required.")
		fmt.Println("example: task-cli add 1")
		os.Exit(1)
	}

	actionName := args[0]

	switch actionName {
	case "add":
		if len(args) < 2 {
			fmt.Println("add command requires a task name..")
			return
		}
		// add
	case "update":
		if len(args) < 3 {
			fmt.Println("update command requires an index and a task name..")
			return
		}

		taskID := args[2]
		validateID(taskID)

		// update
	case "delete":
		if len(args) < 2 {
			fmt.Println("delete command requires an index..")
			return
		}

		taskID := args[2]
		validateID(taskID)

		// delete
	case "mark-in-progress":
		if len(args) < 2 {
			fmt.Println("mark-in-progress command requires an index..")
			return
		}

		taskID := args[2]
		validateID(taskID)

		// mark-in-progress
	case "mark-done":
		if len(args) < 2 {
			fmt.Println("mark-done command requires an index..")
		}

		taskID := args[2]
		validateID(taskID)

		//mark-done
	case "list":
		if len(args) < 2 {
			//list all tasks
		}

		taskID := args[2]
		validateID(taskID)

		// list (done, todo, in-progress)
	}
}

func validateID(id string) {
	_, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("task id must be an integer..")
		return
	}
}
