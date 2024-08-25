package repository

import (
	"encoding/json"
	"errors"
	"os"
	"task-tracker/internal/config"
	"task-tracker/internal/models"
	"task-tracker/internal/utils"
	"time"
)

type Repository interface {
	GetAllTasks() ([]*models.Task, error)
	AddTask(task models.Task) error
	DeleteTask(id int) error
	UpdateTask(id int, description, status string) error
}

type jsonRepository struct{}

func NewJSONRepository() Repository {
	return jsonRepository{}
}

func (j jsonRepository) GetAllTasks() ([]*models.Task, error) {
	f, err := os.ReadFile(config.Filename)
	if err != nil {
		return nil, err
	}

	var tasks []*models.Task
	if err := json.Unmarshal(f, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r jsonRepository) AddTask(task models.Task) error {
	tasks, err := r.GetAllTasks()
	if err != nil {
		return err
	}

	tasks = append(tasks, &task)
	return utils.WriteJSON(tasks)
}

func (r jsonRepository) DeleteTask(id int) error {
	tasks, err := r.GetAllTasks()
	if err != nil {
		return err
	}

	taskIndex, ok := findTaskIndex(tasks, id)
	if !ok {
		return errors.New("task not found")
	}

	tasks = append(tasks[:taskIndex], tasks[taskIndex+1:]...)
	if err := utils.WriteJSON(tasks); err != nil {
		return err
	}

	return nil
}

func findTaskIndex(tasks []*models.Task, id int) (int, bool) {
	var index int
	var ok = false
	for i, t := range tasks {
		if t.Id == id {
			index, ok = i, true
			break
		}
	}
	return index, ok
}

func (r jsonRepository) UpdateTask(id int, description string, status string) error {
	tasks, err := r.GetAllTasks()
	if err != nil {
		return err
	}

	var task *models.Task
	for _, t := range tasks {
		if t.Id == id {
			task = t
		}
	}

	if task == nil {
		return errors.New("task not found")
	}

	if description != "" {
		task.Description = description
	}

	if status != "" {
		task.Status = status
	}

	task.UpdatedAt = time.Now().String()

	if err := utils.WriteJSON(tasks); err != nil {
		return err
	}

	return nil
}
