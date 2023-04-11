package data

import (
	userV1 "fkratos/api/rpc_user/v1"
	"fkratos/app/bff_admin/internal/conf"
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
	// TODO wrapped database client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{}, cleanup, nil
}

func NewUserServiceClient(r registry.Discovery, c *conf.Bootstrap) userV1.UserClient {
	return userV1.NewUserClient(nil)
}
