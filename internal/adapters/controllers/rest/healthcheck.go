package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleHealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World Indie-developers!")
}
