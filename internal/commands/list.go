package commands

import (
	"fmt"
	"task-tracker/internal/models"
	"task-tracker/internal/repository"
)

func List(args []string, repo repository.Repository) (string, error) {
	tasks, err := repo.GetAllTasks()
	if err != nil {
		return "", err
	}

	if len(args) > 0 {
		status := args[0]
		filteredTasks := []*models.Task{}

		for _, task := range tasks {
			if task.Status == status {
				filteredTasks = append(filteredTasks, task)
			}
		}

		tasks = filteredTasks
	}

	var result string
	for _, task := range tasks {
		result += fmt.Sprintf("%d - %s - %s\n", task.Id, task.Description, task.Status)
	}

	return result, nil
}
