package user

import (
	"context"

	"gitlab.com/indie-developers/go-api-echo-template/internal/domains/models"
)

func (u *User) GetAll(ctx context.Context) ([]models.User, error) {
	return u.db.FindAll(ctx)
}
