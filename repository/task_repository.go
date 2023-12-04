package repository

import "gorm.io/gorm"

type taskRepository struct {
	database *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{
		database: db,
	}
}

type TaskRepository interface {
}
