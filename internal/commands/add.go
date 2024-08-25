package commands

import (
	"errors"
	"fmt"
	"task-tracker/internal/models"
	"task-tracker/internal/repository"
)

func AddTask(args []string, repo repository.Repository) (string, error) {
	if len(args) == 0 {
		return "", errors.New("task name is missing")
	}

	tasks, err := repo.GetAllTasks()
	if err != nil {
		return "", nil
	}

	description := args[0]
	for _, task := range tasks {
		if task.Description == description {
			return "", errors.New("this task already exists")
		}
	}

	id := makeID(tasks)
	task := models.NewTask(id, description)

	if err := repo.AddTask(task); err != nil {
		return "", err
	}

	return fmt.Sprintf("Task added successfully (ID: %d)\n", task.Id), nil
}

func makeID(tasks []*models.Task) int {
	return tasks[len(tasks)-1].Id + 1
}
