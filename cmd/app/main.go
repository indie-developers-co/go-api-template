package main

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/indie-developers/go-api-echo-template/api/controllers"
)

func main() {
	e := echo.New()
	e.GET("/healthcheck", controllers.HealthCheck)

	e.Logger.Fatal(e.Start(":8080"))
}
