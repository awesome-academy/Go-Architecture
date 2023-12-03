package repository

import (
	"Go-Architecture/domain/entity"
	"context"
	"gorm.io/gorm"
)

type userRepository struct {
	database *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{database: db}
}

type UserRepository interface {
	GetByEmail(ctx context.Context, email string) (entity.User, error)
	Create(ctx context.Context, user *entity.User) error
	GetByID(ctx context.Context, id string) (entity.User, error)
}

func (ur userRepository) GetByEmail(ctx context.Context, email string) (entity.User, error) {
	var user entity.User
	db := ur.database.Where("email = ?", email).Find(&user)
	if db.Error != nil {
		return entity.User{}, db.Error
	}
	return user, nil
}

func (ur userRepository) Create(ctx context.Context, user *entity.User) error {
	if err := ur.database.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (ur userRepository) GetByID(ctx context.Context, email string) (entity.User, error) {
	var user entity.User
	db := ur.database.Where("id = ?", email).Find(&user)
	if db.Error != nil {
		return entity.User{}, db.Error
	}
	return user, nil
}
