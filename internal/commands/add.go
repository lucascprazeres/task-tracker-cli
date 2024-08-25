package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"task-tracker/internal/config"
	"task-tracker/internal/models"
	"task-tracker/internal/utils"
)

func AddTask(args []string) (string, error) {
	if len(args) == 0 {
		return "", errors.New("task name is missing")
	}

	f, err := os.ReadFile(config.Filename)
	if err != nil {
		return "", err
	}

	var tasks []models.Task
	if err := json.Unmarshal(f, &tasks); err != nil {
		return "", err
	}

	description := args[0]
	for _, task := range tasks {
		if task.Description == description {
			return "", errors.New("this task already exists")
		}
	}

	id := tasks[len(tasks)-1].Id + 1
	task := models.NewTask(id, description)
	tasks = append(tasks, task)

	if err := utils.WriteJSON(tasks); err != nil {
		return "", err
	}

	return fmt.Sprintf("Task added successfully (ID: %d)\n", task.Id), nil
}
