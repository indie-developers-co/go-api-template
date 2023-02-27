package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/gommon/log"
	"gitlab.com/indie-developers/go-api-echo-template/cmd/container"
	"gitlab.com/indie-developers/go-api-echo-template/pkg/migrations"
	"gitlab.com/indie-developers/go-api-echo-template/pkg/router"
	"gitlab.com/indie-developers/go-api-echo-template/pkg/server"
	"gitlab.com/indie-developers/go-api-echo-template/pkg/server/rpc"
	"gitlab.com/indie-developers/go-api-echo-template/pkg/validator"
)

//	@title			Go API Template
//	@version		0.1.0
//	@description	this is a template created with hexagonal architecture for Golang APIs, it includes multiple cases using popular web frameworks like Gin or Echo, and gRPC. For more information contact us to our email.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	indie-developers
//	@contact.email	shorcutbot.indiedevelopers@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@BasePath	/api/v1/template/
func main() {
	c := container.BuildContainer()
	err := c.Invoke(func(migrations migrations.Migrations, server server.Server, requestValidator validator.Validator,
		serverRouter router.Router, rpcServer rpc.RpcServer) {
		migrations.Run()

		errs := make(chan error)

		go func() {
			c := make(chan os.Signal)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
			errs <- fmt.Errorf("%s", <-c)
		}()

		go func() {
			// TODO: add config CORS
			server.Debug(true) // TODO: Validate according to the environment
			server.SetRequestValidator(requestValidator)
			serverRouter.ApplyConfiguration()
			server.Run(":8080") // TODO: handle using env vars
		}()

		go func() {
			rpcServer.Run(":50051")
		}()

		log.Error("exit", <-errs)
	})
	if err != nil {
		panic(err)
	}
}
