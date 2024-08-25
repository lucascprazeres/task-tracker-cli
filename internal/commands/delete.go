package commands

import (
	"errors"
	"fmt"
	"strconv"
	"task-tracker/internal/repository"
)

func Delete(args []string, repo repository.Repository) (string, error) {
	if len(args) == 0 {
		return "", errors.New("missing task ID")
	}

	taskID, err := strconv.Atoi(args[0])
	if err != nil {
		return "", errors.New("invalid ID")
	}

	if err := repo.DeleteTask(taskID); err != nil {
		return "", err
	}

	return fmt.Sprintf("Task removed successfully (ID: %d)\n", taskID), nil
}
