package domain

import (
	"Go-Architecture/domain/entity"
	"context"
)

type TaskUsecase interface {
	Create(c context.Context, task *entity.Task) error
}
