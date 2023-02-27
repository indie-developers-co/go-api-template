package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gitlab.com/indie-developers/go-api-echo-template/internal/domains/models/requests"
	"gitlab.com/indie-developers/go-api-echo-template/internal/domains/models/responses"
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

// CreateUser godoc
//
//	@ID				create-user
//	@Summary		create-user
//	@Tags			User Controller
//	@Description	Create a new user
//	@Param			x-application-id	header		string						true	"requester name"
//	@Param			x-request-id		header		string						false	"UUID request"
//	@Param			request				body		requests.CreateUserRequest	true	"Request"
//	@Success		201					{string}	string						"user has been created successfully"
//	@Failure		500
//	@x-codeSamples	file
//	@Router			/user [post]
func (u *User) CreateUser(c echo.Context) error {
	ctx := c.Request().Context()
	var model requests.CreateUserRequest

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

	return c.JSON(http.StatusCreated, "user has been created successfully")
}

// GetUsers godoc
//
//	@ID				get-users
//	@Summary		get-users
//	@Tags			User Controller
//	@Description	Get all users registered in our database
//	@Param			x-application-id	header	string	true	"requester name"
//	@Param			x-request-id		header	string	false	"UUID request"
//	@Success		200					{array}	responses.GetUsersResponse
//	@Failure		500
//	@x-codeSamples	file
//	@Router			/user [get]
func (u *User) GetUsers(c echo.Context) error {
	ctx := c.Request().Context()

	users, err := u.userUseCases.GetAll(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, responses.ConvertToGetUsersResponseStruct(users))
}
