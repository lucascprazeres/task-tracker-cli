package repository

import (
	"errors"
	"task-tracker/internal/models"
	"time"
)

type mockRepository struct {
	tasks []*models.Task
}

func NewMockRepository() Repository {
	return &mockRepository{
		tasks: make([]*models.Task, 5),
	}
}

func (m *mockRepository) AddTask(task models.Task) error {
	m.tasks = append(m.tasks, &task)
	return nil
}

func (m *mockRepository) DeleteTask(id int) error {
	var index int
	var ok = false
	for i, task := range m.tasks {
		if task.Id == id {
			index = i
			ok = true
			break
		}
	}

	if !ok {
		return errors.New("task not found")
	}

	m.tasks = append(m.tasks[:index], m.tasks[index+1:]...)

	return nil
}

func (m *mockRepository) GetAllTasks() ([]*models.Task, error) {
	return m.tasks, nil
}

func (m mockRepository) UpdateTask(id int, description string, status string) error {
	var task *models.Task
	for _, t := range m.tasks {
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
	return nil
}
