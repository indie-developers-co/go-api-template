package storage

import (
	"context"

	"gitlab.com/indie-developers/go-api-echo-template/internal/domains/entities"
	"gitlab.com/indie-developers/go-api-echo-template/internal/domains/models"
	"gitlab.com/indie-developers/go-api-echo-template/internal/domains/models/request"
	"gitlab.com/indie-developers/go-api-echo-template/internal/domains/repositories"
	"gorm.io/gorm"
)

type User struct {
	client *gorm.DB
}

func NewUser(client *gorm.DB) repositories.UserStorage {
	return &User{client: client}
}

func (u *User) New(ctx context.Context, user request.User) error {
	userEntity := entities.User{
		Name:     user.Name,
		LastName: user.LastName,
		Email:    user.Email,
		IsActive: true,
	}

	result := u.client.WithContext(ctx).Create(&userEntity)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (u *User) FindAll(ctx context.Context) ([]models.User, error) {
	var users []models.User
	result := u.client.WithContext(ctx).
		Select("name", "last_name", "email", "is_active").
		Where("deleted_at is null").
		Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
