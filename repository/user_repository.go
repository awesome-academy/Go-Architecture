package repository

import (
	"Go-Architecture/domain/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetByEmail(ctx interface{}, email string) (entity.User, error)
}

type userRepository struct {
	database *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{database: db}
}

func (ur userRepository) GetByEmail(ctx interface{}, email string) (entity.User, error) {
	var user entity.User
	db := ur.database.Where("email = ?", email).Find(&user)
	if db.Error != nil {
		return entity.User{}, db.Error
	}
	return user, nil
}
