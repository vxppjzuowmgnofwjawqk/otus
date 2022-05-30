package models

type Task struct {
	Id   string `json:"id"`
	Body string `json:"body"`
}

type TaskList struct {
	Tasks []Task `json:"tasks"`
}
