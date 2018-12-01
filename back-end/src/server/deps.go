package main

import (
	api "server/animus/apis"
	"go.uber.org/dig"
	"server/animus/config"
	"server/animus/services"
	"server/animus/utils"

)

func BuildContainer() *dig.Container {
	container := dig.New()

	container.Provide(config.NewAppConfig)
	container.Provide(utils.NewApiUtils)
	container.Provide(services.NewAccountService)
	container.Provide(api.NewAccountApi)
	container.Provide(NewServer)

	return container
}
