package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"task-tracker/internal/config"
	"task-tracker/internal/models"
	"task-tracker/internal/utils"
)

func Delete(args []string) (string, error) {
	if len(args) == 0 {
		return "", errors.New("missing task ID")
	}

	taskID, err := strconv.Atoi(args[0])
	if err != nil {
		return "", errors.New("invalid ID")
	}

	f, err := os.ReadFile(config.Filename)
	if err != nil {
		return "", err
	}

	var tasks []*models.Task
	if err := json.Unmarshal(f, &tasks); err != nil {
		return "", err
	}

	var taskIndex int
	var task *models.Task
	for i, t := range tasks {
		if t.Id == taskID {
			taskIndex = i
			task = t
		}
	}

	if task == nil {
		return "", errors.New("task not found")
	}

	tasks = append(tasks[:taskIndex], tasks[taskIndex+1:]...)
	if err := utils.WriteJSON(tasks); err != nil {
		return "", err
	}

	return fmt.Sprintf("Task removed successfully (ID: %d)\n", taskID), nil
}
