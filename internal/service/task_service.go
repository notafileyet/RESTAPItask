package service

import (
	"APIhendler/internal/repository"
)

type TaskService struct {
	Repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{Repo: repo}
}

func (s *TaskService) CreateTask(task *repository.Task) error {
	return s.Repo.Create(task)
}

func (s *TaskService) GetAllTasks() ([]repository.Task, error) {
	return s.Repo.GetAll()
}

func (s *TaskService) GetTaskByID(id uint) (*repository.Task, error) {
	return s.Repo.GetByID(id)
}

func (s *TaskService) UpdateTask(task *repository.Task) error {
	return s.Repo.Update(task)
}

func (s *TaskService) DeleteTask(id uint) error {
	return s.Repo.Delete(id)
}
