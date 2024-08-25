package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"task-tracker/internal/config"
	"task-tracker/internal/models"
)

var availableStatus = []string{"todo", "in-progress", "done"}

func List(args []string) (string, error) {
	f, err := os.ReadFile(config.Filename)
	if err != nil {
		return "", err
	}

	var tasks []models.Task
	if err := json.Unmarshal(f, &tasks); err != nil {
		return "", err
	}

	if len(args) > 0 && slices.Contains(availableStatus, args[0]) {
		filteredTasks := []models.Task{}
		for _, task := range tasks {
			if task.Status == args[0] {
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
