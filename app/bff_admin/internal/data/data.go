package data

import (
	"context"
	sysV1 "fkratos/api/rpc_sys/v1"
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
		log:        l,
		userClient: userClient,
	}
	cleanup := func() {
		l.Info("closing the data resources")
	}
	return d, cleanup, nil
}

func NewUserServiceClient(c *conf.Bootstrap, r registry.Discovery) userV1.UserClient {
	return userV1.NewUserClient(bootstrap.NewGrpcClient(context.Background(), r, c.Registry.Type, constant.RpcUser, c.Server.Grpc.GetTimeout()))
}

func NewSysAdminServiceClient(c *conf.Bootstrap, r registry.Discovery) sysV1.AdminClient {
	return sysV1.NewAdminClient(bootstrap.NewGrpcClient(context.Background(), r, c.Registry.Type, constant.RpcSys, c.Server.Grpc.GetTimeout()))
}
