//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"fkratos/app/bff_api/internal/server"
	"fkratos/app/bff_api/internal/service"
	"fkratos/internal/bootstrap/conf"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Bootstrap, log.Logger, registry.Registrar) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, service.ProviderSet, newApp))
}
