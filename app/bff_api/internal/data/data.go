package data

import (
	"context"
	userV1 "fkratos/api/rpc_user/v1"
	"fkratos/internal/bootstrap"
	"fkratos/internal/bootstrap/conf"
	"fkratos/internal/constant"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
)

// Data .
type Data struct {
}

// NewData .
func NewData(c *conf.Bootstrap, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{}, cleanup, nil
}

func NewUserServiceClient(c *conf.Bootstrap, r registry.Discovery) userV1.UserClient {
	return userV1.NewUserClient(bootstrap.NewGrpcClient(context.Background(), r, c.Registry.Type, constant.RpcUser, c.Server.Grpc.GetTimeout()))
}
