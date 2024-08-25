package commands

import (
	"errors"
	"fmt"
	"strconv"
	"task-tracker/internal/repository"
)

func MarkDone(args []string, repo repository.Repository) (string, error) {
	if len(args) == 0 {
		return "", errors.New("missing task ID")
	}

	taskID, err := strconv.Atoi(args[0])
	if err != nil {
		return "", errors.New("invalid ID")
	}

	if err := repo.UpdateTask(taskID, "", "done"); err != nil {
		return "", err
	}

	return fmt.Sprintf("Task updated successfully (ID: %d)\n", taskID), nil
}
