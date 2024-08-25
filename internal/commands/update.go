package commands

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"task-tracker/internal/config"
	"task-tracker/internal/models"
	"task-tracker/internal/utils"
	"time"
)

func Update(args []string) (string, error) {
	if len(args) < 2 {
		return "", errors.New("task ID and description required")
	}

	taskID, err := strconv.Atoi(args[0])
	if err != nil {
		return "", errors.New("invalid ID")
	}

	description := args[1]
	if len(description) == 0 {
		return "", errors.New("description is empty")
	}

	f, err := os.ReadFile(config.Filename)
	if err != nil {
		return "", err
	}

	var tasks []*models.Task
	if err := json.Unmarshal(f, &tasks); err != nil {
		return "", err
	}

	var task *models.Task
	for _, t := range tasks {
		if t.Id == taskID {
			task = t
		}
	}

	if task == nil {
		return "", errors.New("task not found")
	}

	task.Description = description
	task.UpdatedAt = time.Now().String()

	if err := utils.WriteJSON(tasks); err != nil {
		return "", err
	}

	return "", nil
}
