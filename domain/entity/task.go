package entity

type Task struct {
	BaseEntity

	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description"`
	UserID      int    `json:"userID"`
}

func (Task) TableName() string {
	return "tasks"
}
