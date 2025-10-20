package tasksRepo

import (
	"APIhendler/internal/userService/orm"
	"time"
)

type Task struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	UserID uint     `json:"user_id" gorm:"not null"`
	User   orm.User `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

type TaskRepositoryInterface interface {
	Create(task *Task) error
	GetAll() ([]Task, error)
	GetByID(id uint) (*Task, error)
	Update(task *Task) error
	Delete(id uint) error
	GetTasksByUserID(userID uint) ([]Task, error)
}

var _ TaskRepositoryInterface = &TaskRepository{}
