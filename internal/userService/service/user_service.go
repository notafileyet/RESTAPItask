package service

import (
	"APIhendler/internal/tasksRepo"
	"APIhendler/internal/userService/orm"
	"APIhendler/internal/userService/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserService struct {
	Repo     repository.UserInterface
	TaskRepo tasksRepo.TaskRepositoryInterface
}

func NewUserService(repo repository.UserInterface, taskRepo tasksRepo.TaskRepositoryInterface) *UserService {
	return &UserService{Repo: repo, TaskRepo: taskRepo}
}

func (s *UserService) CreateUser(user *orm.User) error {
	return s.Repo.Create(user)
}

func (s *UserService) GetAllUsers() ([]orm.User, error) {
	return s.Repo.GetAll()
}

func (s *UserService) GetUserByID(id uint) (*orm.User, error) {
	return s.Repo.GetByID(id)
}

func (s *UserService) UpdateUser(user *orm.User) error {
	return s.Repo.Update(user)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.Repo.Delete(id)
}

func (s *UserService) GetTasksForUser(userID uint) ([]tasksRepo.Task, error) {
	_, err := s.Repo.GetByID(userID)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	return s.TaskRepo.GetTasksByUserID(userID)
}

func (s *UserService) GetUserWithTasks(userID uint) (*orm.User, error) {
	return s.Repo.GetUserWithTasks(userID)
}
