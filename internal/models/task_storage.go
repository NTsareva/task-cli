package models

type TaskStorage struct {
	ID     int        `json:"id"`
	Name   string     `json:"name"`
	Status TaskStatus `json:"status"`
}
