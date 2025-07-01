package repository

import (
	"gorm.io/gorm"
)

type Task struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

type TaskRepository struct {
	DB *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	db.AutoMigrate(&Task{})
	return &TaskRepository{DB: db}
}

func (r *TaskRepository) Create(task *Task) error {
	return r.DB.Create(task).Error
}

func (r *TaskRepository) GetAll() ([]Task, error) {
	var tasks []Task
	err := r.DB.Find(&tasks).Error
	return tasks, err
}

func (r *TaskRepository) GetByID(id uint) (*Task, error) {
	var task Task
	err := r.DB.First(&task, id).Error
	return &task, err
}

func (r *TaskRepository) Update(task *Task) error {
	return r.DB.Save(task).Error
}

func (r *TaskRepository) Delete(id uint) error {
	return r.DB.Delete(&Task{}, id).Error
}
