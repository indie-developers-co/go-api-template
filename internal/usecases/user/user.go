package user

import (
	"gitlab.com/indie-developers/go-api-echo-template/internal/domains/repositories"
)

type User struct {
	db repositories.UserStorage
}

func NewUser(db repositories.UserStorage) repositories.UserUseCases {
	return &User{db: db}
}
