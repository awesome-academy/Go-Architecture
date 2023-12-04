package repository

import (
	"Go-Architecture/domain/entity"
	"context"
	"gorm.io/gorm"
)

type taskRepository struct {
	database *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{
		database: db,
	}
}

type TaskRepository interface {
	Create(ctx context.Context, task *entity.Task) error
}

func (tr taskRepository) Create(ctx context.Context, task *entity.Task) error {
	if err := tr.database.Create(&task).Error; err != nil {
		return err
	}
	return nil
}
