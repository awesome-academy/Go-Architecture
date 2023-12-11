package entity

import "gorm.io/gorm"

type Task struct {
	gorm.Model

	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description"`
	UserID      string `json:"userID"`
}

func (Task) TableName() string {
	return "tasks"
}
