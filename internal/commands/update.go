package commands

import (
	"errors"
	"fmt"
	"strconv"
	"task-tracker/internal/repository"
)

func Update(args []string, repo repository.Repository) (string, error) {
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

	if err := repo.UpdateTask(taskID, description, ""); err != nil {
		return "", err
	}

	return fmt.Sprintf("Task updated successfully (ID: %d)\n", taskID), nil
}
