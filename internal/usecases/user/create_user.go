package user

import (
	"context"

	"gitlab.com/indie-developers/go-api-echo-template/internal/domains/models/requests"
)

func (u *User) Create(ctx context.Context, user requests.CreateUserRequest) error {
	return u.db.New(ctx, user)
}
