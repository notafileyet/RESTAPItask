package tasksService

import (
	"APIhendler/internal/tasksRepo"
)

type TaskService struct {
	Repo tasksRepo.TaskRepositoryInterface
}

func NewTaskService(repo tasksRepo.TaskRepositoryInterface) *TaskService {
	return &TaskService{Repo: repo}
}

func (s *TaskService) CreateTask(task *tasksRepo.Task) error {
	return s.Repo.Create(task)
}

func (s *TaskService) GetAllTasks() ([]tasksRepo.Task, error) {
	return s.Repo.GetAll()
}

func (s *TaskService) GetTaskByID(id uint) (*tasksRepo.Task, error) {
	return s.Repo.GetByID(id)
}

func (s *TaskService) UpdateTask(task *tasksRepo.Task) error {
	return s.Repo.Update(task)
}

func (s *TaskService) DeleteTask(id uint) error {
	return s.Repo.Delete(id)
}

func (s *TaskService) GetTasksByUserID(userID uint) ([]tasksRepo.Task, error) {
	return s.Repo.GetTasksByUserID(userID)
}
