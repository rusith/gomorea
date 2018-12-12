package main

import (
	"go.uber.org/dig"
	api "server/animus/apis"
	"server/animus/config"
	"server/animus/providers"
	"server/animus/services"
	"server/animus/utils"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	container.Provide(providers.NewDatabaseProvider)
	container.Provide(providers.NewUserProvider)
	container.Provide(config.NewAppConfig)
	container.Provide(utils.NewApiUtils)
	container.Provide(services.NewAccountService)
	container.Provide(api.NewAccountApi)
	container.Provide(NewServer)

	return container
}
