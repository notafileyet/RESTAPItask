package mocks

import (
	"APIhendler/internal/tasksRepo"
)

type MockTaskRepository struct {
	CreateFunc  func(task *tasksRepo.Task) error
	GetAllFunc  func() ([]tasksRepo.Task, error)
	GetByIDFunc func(id uint) (*tasksRepo.Task, error)
	UpdateFunc  func(task *tasksRepo.Task) error
	DeleteFunc  func(id uint) error
}

func (m *MockTaskRepository) Create(task *tasksRepo.Task) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(task)
	}
	return nil
}

func (m *MockTaskRepository) GetAll() ([]tasksRepo.Task, error) {
	if m.GetAllFunc != nil {
		return m.GetAllFunc()
	}
	return []tasksRepo.Task{}, nil
}

func (m *MockTaskRepository) GetByID(id uint) (*tasksRepo.Task, error) {
	if m.GetByIDFunc != nil {
		return m.GetByIDFunc(id)
	}
	return &tasksRepo.Task{}, nil
}

func (m *MockTaskRepository) Update(task *tasksRepo.Task) error {
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
