package tasksRepo

import (
	"gorm.io/gorm"
)

type TaskRepository struct {
	DB *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{DB: db}
}

func (r *TaskRepository) Create(task *Task) error {
	return r.DB.Create(task).Error
}

func (r *TaskRepository) GetAll() ([]Task, error) {
	var tasks []Task
	err := r.DB.Preload("User").Find(&tasks).Error
	return tasks, err
}

func (r *TaskRepository) GetByID(id uint) (*Task, error) {
	var task Task
	err := r.DB.Preload("User").First(&task, id).Error
	return &task, err
}

func (r *TaskRepository) Update(task *Task) error {
	return r.DB.Save(task).Error
}

func (r *TaskRepository) Delete(id uint) error {
	return r.DB.Delete(&Task{}, id).Error
}

func (r *TaskRepository) GetTasksByUserID(userID uint) ([]Task, error) {
	var tasks []Task
	err := r.DB.Where("user_id = ?", userID).Preload("User").Find(&tasks).Error
	return tasks, err
}
