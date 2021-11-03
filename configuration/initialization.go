package configuration

import (
	"go-clean-arch/core/businesslogic"
	"go-clean-arch/persistence"
	"go-clean-arch/web"
)

func WireServer() web.Server {
	return web.NewServer(
		businesslogic.NewUserUseCases(persistence.NewUserRepository(
			persistence.DBConfig{
				Address: "http://localhost:4566",
				Region:  "eu-west-1",
				Profile: "localstack",
				ID:      "test",
				Secret:  "test",
			})))
}
