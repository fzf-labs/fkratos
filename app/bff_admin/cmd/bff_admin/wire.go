//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"fkratos/app/bff_admin/internal/data"
	"fkratos/app/bff_admin/internal/server"
	"fkratos/app/bff_admin/internal/service"
	"fkratos/internal/bootstrap/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
)

// wireApp init kratos application.
func wireApp(*conf.Bootstrap, log.Logger, registry.Registrar, registry.Discovery) (*kratos.App, func(), error) {
	panic(wire.Build(data.ProviderSet, server.ProviderSet, service.ProviderSet, newApp))
}
