package models

import "time"

type Task struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

func NewTask(id int, description string) Task {
	return Task{
		Id:          id,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now().String(),
		UpdatedAt:   time.Now().String(),
	}
}
