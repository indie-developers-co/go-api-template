package container

import (
	"log"

	"github.com/labstack/echo/v4"
	"gitlab.com/indie-developers/go-api-echo-template/internal/adapters/controllers/rest"
	"gitlab.com/indie-developers/go-api-echo-template/internal/adapters/controllers/rpc"
	"gitlab.com/indie-developers/go-api-echo-template/internal/adapters/storage"
	"gitlab.com/indie-developers/go-api-echo-template/internal/domains/entities"
	"gitlab.com/indie-developers/go-api-echo-template/internal/usecases/user"
	"gitlab.com/indie-developers/go-api-echo-template/pkg/migrations/gorm_migrations"
	"gitlab.com/indie-developers/go-api-echo-template/pkg/router/echorouter"
	"gitlab.com/indie-developers/go-api-echo-template/pkg/server/grpcserver"
	"gitlab.com/indie-developers/go-api-echo-template/pkg/server/rest/echoserver"
	"gitlab.com/indie-developers/go-api-echo-template/pkg/storage/postgres"
	"gitlab.com/indie-developers/go-api-echo-template/pkg/validator/govalidator"
	"go.uber.org/dig"
	"google.golang.org/grpc"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	// server settings (api framework)
	checkError(container.Provide(echo.New))
	checkError(container.Provide(echoserver.NewServer))
	checkError(container.Provide(govalidator.NewRequestValidator))

	// databases settings and migrations
	checkError(container.Provide(postgres.Connection))
	checkError(container.Provide(entities.Provide))
	checkError(container.Provide(gorm_migrations.NewMigrator))

	// user domains
	checkError(container.Provide(storage.NewUser))
	checkError(container.Provide(user.NewUser))
	checkError(container.Provide(rest.NewUser))
	checkError(container.Provide(rpc.NewUserController))

	// http router
	checkError(container.Provide(echorouter.NewRouter))

	// grpc framework
	checkError(container.Provide(grpc.NewServer))
	checkError(container.Provide(grpcserver.NewServer))

	return container
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("error injecting dependencies: %s", err.Error())
	}
}
