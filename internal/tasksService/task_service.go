package tasksService

import (
	"APIhendler/internal/tasksRepo"
	userServiceRepo "APIhendler/internal/userService/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TaskService struct {
	Repo     tasksRepo.TaskRepositoryInterface
	UserRepo userServiceRepo.UserInterface
}

func NewTaskService(repo tasksRepo.TaskRepositoryInterface, userRepo userServiceRepo.UserInterface) *TaskService {
	return &TaskService{Repo: repo, UserRepo: userRepo}
}

func (s *TaskService) CreateTask(task *tasksRepo.Task) error {
	_, err := s.UserRepo.GetByID(task.UserID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

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
