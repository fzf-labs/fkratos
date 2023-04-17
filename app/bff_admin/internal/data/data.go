package data

import (
	"context"
	userV1 "fkratos/api/rpc_user/v1"
	"fkratos/internal/bootstrap"
	"fkratos/internal/bootstrap/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewUserServiceClient,
)

// Data .
type Data struct {
	log        *log.Helper
	userClient userV1.UserClient
}

// NewData .
func NewData(c *conf.Bootstrap, logger log.Logger, userClient userV1.UserClient) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "data/bff-admin"))
	d := &Data{
		userClient: userClient,
	}
	cleanup := func() {
		l.Info("closing the data resources")
	}
	return d, cleanup, nil
}

func NewUserServiceClient(r registry.Discovery, c *conf.Bootstrap) userV1.UserClient {
	return userV1.NewUserClient(bootstrap.NewGrpcClient(context.Background(), r, c.GetServiceName(), c.Server.Grpc.GetTimeout()))
}
