package usecase

import (
	"Go-Architecture/domain"
	"Go-Architecture/repository"
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
