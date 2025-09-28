package tasksService

import (
	"errors"
	"testing"

	"APIhendler/internal/tasksRepo"
	"APIhendler/internal/tasksService/mocks"

	"github.com/stretchr/testify/assert"
)

func TestTaskService_CreateTask(t *testing.T) {
	t.Run("successful creation", func(t *testing.T) {
		mockRepo := &mocks.MockTaskRepository{
			CreateFunc: func(task *tasksRepo.Task) error {
				return nil
			},
		}

		service := NewTaskService(mockRepo)
		task := &tasksRepo.Task{Title: "Test Task", Status: "todo"}

		err := service.CreateTask(task)
		assert.NoError(t, err)
	})

	t.Run("creation error", func(t *testing.T) {
		mockRepo := &mocks.MockTaskRepository{
			CreateFunc: func(task *tasksRepo.Task) error {
				return errors.New("database error")
			},
		}

		service := NewTaskService(mockRepo)
		task := &tasksRepo.Task{Title: "Test Task", Status: "todo"}

		err := service.CreateTask(task)
		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
	})
}

func TestTaskService_GetAllTasks(t *testing.T) {
	t.Run("successful retrieval", func(t *testing.T) {
		expectedTasks := []tasksRepo.Task{
			{ID: 1, Title: "Task 1", Status: "todo"},
			{ID: 2, Title: "Task 2", Status: "done"},
		}

		mockRepo := &mocks.MockTaskRepository{
			GetAllFunc: func() ([]tasksRepo.Task, error) {
				return expectedTasks, nil
			},
		}

		service := NewTaskService(mockRepo)

		tasks, err := service.GetAllTasks()
		assert.NoError(t, err)
		assert.Equal(t, expectedTasks, tasks)
	})

	t.Run("retrieval error", func(t *testing.T) {
		mockRepo := &mocks.MockTaskRepository{
			GetAllFunc: func() ([]tasksRepo.Task, error) {
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
		expectedTask := &tasksRepo.Task{ID: 1, Title: "Test Task", Status: "todo"}

		mockRepo := &mocks.MockTaskRepository{
			GetByIDFunc: func(id uint) (*tasksRepo.Task, error) {
				return expectedTask, nil
			},
		}

		service := NewTaskService(mockRepo)

		task, err := service.GetTaskByID(1)
		assert.NoError(t, err)
		assert.Equal(t, expectedTask, task)
	})

	t.Run("not found error", func(t *testing.T) {
		mockRepo := &mocks.MockTaskRepository{
			GetByIDFunc: func(id uint) (*tasksRepo.Task, error) {
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
		mockRepo := &mocks.MockTaskRepository{
			UpdateFunc: func(task *tasksRepo.Task) error {
				return nil
			},
		}

		service := NewTaskService(mockRepo)
		task := &tasksRepo.Task{ID: 1, Title: "Updated Task", Status: "done"}

		err := service.UpdateTask(task)
		assert.NoError(t, err)
	})

	t.Run("update error", func(t *testing.T) {
		mockRepo := &mocks.MockTaskRepository{
			UpdateFunc: func(task *tasksRepo.Task) error {
				return errors.New("update failed")
			},
		}

		service := NewTaskService(mockRepo)
		task := &tasksRepo.Task{ID: 1, Title: "Updated Task", Status: "done"}

		err := service.UpdateTask(task)
		assert.Error(t, err)
		assert.Equal(t, "update failed", err.Error())
	})
}

func TestTaskService_DeleteTask(t *testing.T) {
	t.Run("successful deletion", func(t *testing.T) {
		mockRepo := &mocks.MockTaskRepository{
			DeleteFunc: func(id uint) error {
				return nil
			},
		}

		service := NewTaskService(mockRepo)

		err := service.DeleteTask(1)
		assert.NoError(t, err)
	})

	t.Run("deletion error", func(t *testing.T) {
		mockRepo := &mocks.MockTaskRepository{
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
