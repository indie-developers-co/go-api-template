package main

import (
	"gitlab.com/indie-developers/go-api-echo-template/cmd/container"
	"gitlab.com/indie-developers/go-api-echo-template/pkg/migrations"
	"gitlab.com/indie-developers/go-api-echo-template/pkg/router"
	"gitlab.com/indie-developers/go-api-echo-template/pkg/server"
	"gitlab.com/indie-developers/go-api-echo-template/pkg/validator"
)

func main() {
	c := container.BuildContainer()
	err := c.Invoke(func(migrations migrations.Migrations, server server.Server, requestValidator validator.Validator, serverRouter router.Router) {
		migrations.Run()

		// TODO: add config CORS
		server.Debug(true) // TODO: Validate according to the environment
		server.SetRequestValidator(requestValidator)
		serverRouter.ApplyConfiguration()
		server.Run(":8080") // TODO: handle using env vars
	})
	if err != nil {
		panic(err)
	}
}
