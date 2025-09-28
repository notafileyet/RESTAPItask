package service

import (
	"APIhendler/internal/userService/orm"
	"APIhendler/internal/userService/repository"
)

type UserService struct {
	Repo repository.UserInterface
}

func NewUserService(repo repository.UserInterface) *UserService {
	return &UserService{Repo: repo}
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
