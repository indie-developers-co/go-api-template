package user

import (
	"context"

	"gitlab.com/indie-developers/go-api-echo-template/internal/domains/models/request"
)

func (u *User) Create(ctx context.Context, user request.CreateUserRequest) error {
	return u.db.New(ctx, user)
}
