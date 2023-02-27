package echorouter

import (
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "gitlab.com/indie-developers/go-api-echo-template/docs"

	"gitlab.com/indie-developers/go-api-echo-template/internal/adapters/controllers/rest"
	"gitlab.com/indie-developers/go-api-echo-template/pkg/router"
)

type routers struct {
	server *echo.Echo
	user   rest.UserController
}

func NewRouter(server *echo.Echo, user rest.UserController) router.Router {
	return &routers{
		server: server,
		user:   user,
	}
}

func (r *routers) ApplyConfiguration() {
	r.server.Use(echoMiddleware.RecoverWithConfig(echoMiddleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KiB
		LogLevel:  log.ERROR,
	}))

	r.server.GET("/docs/*", echoSwagger.WrapHandler)

	apiBase := r.server.Group(router.BasePathV1)
	apiBase.GET("/healthcheck", rest.HandleHealthCheck)

	user := apiBase.Group("/user")
	user.POST("", r.user.CreateUser)
	user.GET("", r.user.GetUsers)
}
