package repositories

import (
	"context"

	"gitlab.com/indie-developers/go-api-echo-template/internal/domains/models"
	"gitlab.com/indie-developers/go-api-echo-template/internal/domains/models/request"
)

//go:generate mockery --name UserStorage --output ../../../tests/mocks/
type UserStorage interface {
	New(ctx context.Context, user request.CreateUserRequest) error
	FindAll(ctx context.Context) ([]models.User, error)
}

//go:generate mockery --name UserUseCases --output ../../../tests/mocks/
type UserUseCases interface {
	Create(ctx context.Context, user request.CreateUserRequest) error
	GetAll(ctx context.Context) ([]models.User, error)
}
