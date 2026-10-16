package models

type TaskStatus string

const (
	TODO        TaskStatus = "todo"
	IN_PROGRESS TaskStatus = "in-progress"
	DONE        TaskStatus = "done"
)

func (t TaskStatus) IsValid() bool {
	switch t {
	case TODO, IN_PROGRESS, DONE:
		return true
	default:
		return false
	}
}
