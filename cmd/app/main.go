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
	"gitlab.com/indie-developers/go-api-echo-template/pkg/validator"
)

func main() {
	c := container.BuildContainer()
	err := c.Invoke(func(migrations migrations.Migrations, server server.Server, requestValidator validator.Validator,
		serverRouter router.Router, rpcServer server.RpcServer) {
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
