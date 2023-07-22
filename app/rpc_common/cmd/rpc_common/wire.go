//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"fkratos/app/rpc_common/internal/biz"
	"fkratos/app/rpc_common/internal/data"
	"fkratos/app/rpc_common/internal/server"
	"fkratos/app/rpc_common/internal/service"
	"fkratos/internal/bootstrap/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
)

// wireApp init kratos application.
func wireApp(*conf.Bootstrap, log.Logger, registry.Registrar, registry.Discovery) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, service.ProviderSet, biz.ProviderSet, data.ProviderSet, newApp))
}
