package service

import (
	"APIhendler/internal/repository"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockTaskRepository struct {
	CreateFunc  func(task *repository.Task) error
	GetAllFunc  func() ([]repository.Task, error)
	GetByIDFunc func(id uint) (*repository.Task, error)
	UpdateFunc  func(task *repository.Task) error
	DeleteFunc  func(id uint) error
}

func (m *MockTaskRepository) Create(task *repository.Task) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(task)
	}
	return nil
}

func (m *MockTaskRepository) GetAll() ([]repository.Task, error) {
	if m.GetAllFunc != nil {
		return m.GetAllFunc()
	}
	return []repository.Task{}, nil
}

func (m *MockTaskRepository) GetByID(id uint) (*repository.Task, error) {
	if m.GetByIDFunc != nil {
		return m.GetByIDFunc(id)
	}
	return &repository.Task{}, nil
}

func (m *MockTaskRepository) Update(task *repository.Task) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(task)
	}
	return nil
}

func (m *MockTaskRepository) Delete(id uint) error {
	if m.DeleteFunc != nil {
		return m.DeleteFunc(id)
	}
	return nil
}

func TestTaskService_CreateTask(t *testing.T) {
	t.Run("successful creation", func(t *testing.T) {
		mockRepo := &MockTaskRepository{
			CreateFunc: func(task *repository.Task) error {
				return nil
			},
		}

		service := NewTaskService(mockRepo)
		task := &repository.Task{Title: "Test Task", Status: "todo"}

		err := service.CreateTask(task)
		assert.NoError(t, err)
	})

	t.Run("creation error", func(t *testing.T) {
		mockRepo := &MockTaskRepository{
			CreateFunc: func(task *repository.Task) error {
				return errors.New("database error")
			},
		}

		service := NewTaskService(mockRepo)
		task := &repository.Task{Title: "Test Task", Status: "todo"}

		err := service.CreateTask(task)
		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
	})
}

func TestTaskService_GetAllTasks(t *testing.T) {
	t.Run("successful retrieval", func(t *testing.T) {
		expectedTasks := []repository.Task{
			{ID: 1, Title: "Task 1", Status: "todo"},
			{ID: 2, Title: "Task 2", Status: "done"},
		}

		mockRepo := &MockTaskRepository{
			GetAllFunc: func() ([]repository.Task, error) {
				return expectedTasks, nil
			},
		}

		service := NewTaskService(mockRepo)

		tasks, err := service.GetAllTasks()
		assert.NoError(t, err)
		assert.Equal(t, expectedTasks, tasks)
	})

	t.Run("retrieval error", func(t *testing.T) {
		mockRepo := &MockTaskRepository{
			GetAllFunc: func() ([]repository.Task, error) {
				return nil, errors.New("database error")
			},
		}

		service := NewTaskService(mockRepo)

		tasks, err := service.GetAllTasks()
		assert.Error(t, err)
		assert.Nil(t, tasks)
		assert.Equal(t, "database error", err.Error())
	})
}

func TestTaskService_GetTaskByID(t *testing.T) {
	t.Run("successful retrieval by ID", func(t *testing.T) {
		expectedTask := &repository.Task{ID: 1, Title: "Test Task", Status: "todo"}

		mockRepo := &MockTaskRepository{
			GetByIDFunc: func(id uint) (*repository.Task, error) {
				return expectedTask, nil
			},
		}

		service := NewTaskService(mockRepo)

		task, err := service.GetTaskByID(1)
		assert.NoError(t, err)
		assert.Equal(t, expectedTask, task)
	})

	t.Run("not found error", func(t *testing.T) {
		mockRepo := &MockTaskRepository{
			GetByIDFunc: func(id uint) (*repository.Task, error) {
				return nil, errors.New("task not found")
			},
		}

		service := NewTaskService(mockRepo)

		task, err := service.GetTaskByID(999)
		assert.Error(t, err)
		assert.Nil(t, task)
		assert.Equal(t, "task not found", err.Error())
	})
}

func TestTaskService_UpdateTask(t *testing.T) {
	t.Run("successful update", func(t *testing.T) {
		mockRepo := &MockTaskRepository{
			UpdateFunc: func(task *repository.Task) error {
				return nil
			},
		}

		service := NewTaskService(mockRepo)
		task := &repository.Task{ID: 1, Title: "Updated Task", Status: "done"}

		err := service.UpdateTask(task)
		assert.NoError(t, err)
	})

	t.Run("update error", func(t *testing.T) {
		mockRepo := &MockTaskRepository{
			UpdateFunc: func(task *repository.Task) error {
				return errors.New("update failed")
			},
		}

		service := NewTaskService(mockRepo)
		task := &repository.Task{ID: 1, Title: "Updated Task", Status: "done"}

		err := service.UpdateTask(task)
		assert.Error(t, err)
		assert.Equal(t, "update failed", err.Error())
	})
}

func TestTaskService_DeleteTask(t *testing.T) {
	t.Run("successful deletion", func(t *testing.T) {
		mockRepo := &MockTaskRepository{
			DeleteFunc: func(id uint) error {
				return nil
			},
		}

		service := NewTaskService(mockRepo)

		err := service.DeleteTask(1)
		assert.NoError(t, err)
	})

	t.Run("deletion error", func(t *testing.T) {
		mockRepo := &MockTaskRepository{
			DeleteFunc: func(id uint) error {
				return errors.New("deletion failed")
			},
		}

		service := NewTaskService(mockRepo)

		err := service.DeleteTask(1)
		assert.Error(t, err)
		assert.Equal(t, "deletion failed", err.Error())
	})
}
