package main

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/indie-developers/go-api-echo-template/api/controllers"
)

const baseAPI = "/api/template"

func main() {
	e := echo.New()
	base := e.Group(baseAPI, nil)
	base.GET("/healthcheck", controllers.HealthCheck)

	e.Logger.Fatal(e.Start(":8080"))
}
