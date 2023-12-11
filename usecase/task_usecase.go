package usecase

import (
	"Go-Architecture/domain"
	"Go-Architecture/domain/entity"
	"Go-Architecture/repository"
	"context"
	"time"
)

type taskUsecase struct {
	taskRepository repository.TaskRepository
	contextTimeout time.Duration
}

func NewTaskUsecase(taskRepository repository.TaskRepository, timeout time.Duration) domain.TaskUsecase {
	return &taskUsecase{
		taskRepository: taskRepository,
		contextTimeout: timeout,
	}
}

func (tu taskUsecase) Create(c context.Context, task *entity.Task) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepository.Create(ctx, task)
}
