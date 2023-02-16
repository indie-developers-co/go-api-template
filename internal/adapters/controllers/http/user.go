package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gitlab.com/indie-developers/go-api-echo-template/internal/domains/models/request"
	"gitlab.com/indie-developers/go-api-echo-template/internal/domains/repositories"
)

type UserController interface {
	CreateUser(c echo.Context) error
	GetUsers(c echo.Context) error
}

type User struct {
	userUseCases repositories.UserUseCases
}

func NewUser(useCases repositories.UserUseCases) UserController {
	return &User{userUseCases: useCases}
}

func (u *User) CreateUser(c echo.Context) error {
	ctx := c.Request().Context()
	var model request.User

	if err := c.Bind(&model); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := c.Validate(&model); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := u.userUseCases.Create(ctx, model); err != nil {
		// TODO: create parse function to validate the type of the error
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, nil)
}

func (u *User) GetUsers(c echo.Context) error {
	ctx := c.Request().Context()

	models, err := u.userUseCases.GetAll(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, models)

}
