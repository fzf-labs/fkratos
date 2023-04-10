//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"fkratos/app/rpc/user/internal/biz"
	"fkratos/app/rpc/user/internal/conf"
	"fkratos/app/rpc/user/internal/data"
	"fkratos/app/rpc/user/internal/server"
	"fkratos/app/rpc/user/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, service.ProviderSet, biz.ProviderSet, data.ProviderSet, newApp))
}
