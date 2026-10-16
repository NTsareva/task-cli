package handlers

import "fmt"

func AddHandler(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("add command requires a task name. Example: task-cli add \"create task\"")
	}

	return nil
}

func UpdateHandler(args []string) error {
	if len(args) < 3 {
		return fmt.Errorf("update command requires an index and a task name. Example: task-cli update 1 \"updated task\"")

	}

	return nil
}

func DeleteHandler(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("delete command requires an index. Example: task-cli delete 1")
	}

	return nil
}

func MarkInProgressHandler(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("mark-in-progress command requires an index. Example: task-cli mark-in-progress 1")
	}

	return nil
}

func MarkDoneHandler(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("mark-done command requires an index. Example: task-cli mark-done 1")
	}

	return nil
}

func ListHandler(args []string) error {

	return nil
}
