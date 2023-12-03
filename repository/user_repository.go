package repository

import (
	"gorm.io/gorm"
)

type UserRepository interface {
}

type userRepository struct {
	database *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{database: db}
}
