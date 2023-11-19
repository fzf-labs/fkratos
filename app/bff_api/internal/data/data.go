package data

import (
	"fkratos/app/bff_api/internal/data/rpc"
	"fkratos/internal/bootstrap/conf"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	rpc.NewRPCUserGrpc,
	rpc.NewUserServiceClient,
)

type Data struct {
	log *log.Helper
}

func NewData(c *conf.Bootstrap, logger log.Logger) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", fmt.Sprintf("%s/data", c.Name)))
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		log: l,
	}, cleanup, nil
}
