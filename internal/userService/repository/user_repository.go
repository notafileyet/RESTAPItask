package repository

import (
	"APIhendler/internal/tasksRepo"
	"APIhendler/internal/userService/orm"

	"gorm.io/gorm"
)

type UserInterface interface {
	Create(user *orm.User) error
	GetAll() ([]orm.User, error)
	GetByID(id uint) (*orm.User, error)
	Update(user *orm.User) error
	Delete(id uint) error
	GetUserWithTasks(id uint) (*orm.User, error)
}

type UserRepository struct {
	DB       *gorm.DB
	TaskRepo tasksRepo.TaskRepositoryInterface
}

func NewUserRepository(db *gorm.DB, taskRepo tasksRepo.TaskRepositoryInterface) *UserRepository {
	return &UserRepository{DB: db, TaskRepo: taskRepo}
}

var _ UserInterface = &UserRepository{}

func (r *UserRepository) Create(user *orm.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) GetAll() ([]orm.User, error) {
	var users []orm.User
	err := r.DB.Where("deleted_at IS NULL").Find(&users).Error
	return users, err
}

func (r *UserRepository) GetByID(id uint) (*orm.User, error) {
	var user orm.User
	err := r.DB.Where("id = ? AND deleted_at IS NULL", id).First(&user).Error
	return &user, err
}

func (r *UserRepository) Update(user *orm.User) error {
	return r.DB.Save(user).Error
}

func (r *UserRepository) Delete(id uint) error {
	return r.DB.Delete(&orm.User{}, id).Error
}

func (r *UserRepository) GetUserWithTasks(id uint) (*orm.User, error) {
	var user orm.User
	err := r.DB.Preload("Tasks").Where("id = ? AND deleted_at IS NULL", id).First(&user).Error
	return &user, err
}
