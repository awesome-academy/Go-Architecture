package controller

import (
	"Go-Architecture/domain"
	"github.com/gin-gonic/gin"
)

type TaskController struct {
	TaskUsecase domain.TaskUsecase
}

func (tc TaskController) Create(c *gin.Context) {
	// Bind task
	// Get UserId from header
	// Create task
}

func (tc TaskController) Fetch(c *gin.Context) {
	// TODO implement later
}
