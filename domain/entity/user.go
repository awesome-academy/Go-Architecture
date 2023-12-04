package entity

type User struct {
	BaseEntity

	Name     string `json:"name" gorm:"column:name"`
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"password" gorm:"column:password"`
}

func (User) TableName() string {
	return "users"
}
