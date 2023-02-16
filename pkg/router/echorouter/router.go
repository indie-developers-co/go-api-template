package echorouter

import (
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"gitlab.com/indie-developers/go-api-echo-template/internal/adapters/controllers/http"
	"gitlab.com/indie-developers/go-api-echo-template/pkg/router"
)

type routers struct {
	server *echo.Echo
	user   http.UserController
}

func NewRouter(server *echo.Echo, user http.UserController) router.Router {
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

	apiBase := r.server.Group(router.BasePathV1)
	apiBase.GET("/healthcheck", http.HandleHealthCheck)

	user := apiBase.Group("/user")
	user.POST("", r.user.CreateUser)
	user.GET("", r.user.GetUsers)
}
