package data

import (
	userV1 "fkratos/api/rpc_user/v1"
	"fkratos/internal/rpc"
	"fmt"

	conf "github.com/fzf-labs/fkratos-contrib/api/conf/v1"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	rpc.NewRPCSysGrpc,
	rpc.NewSysAdminServiceClient,
	rpc.NewSysAuthServiceClient,
	rpc.NewSysDashboardServiceClient,
	rpc.NewSysRoleServiceClient,
	rpc.NewSysPermissionServiceClient,
	rpc.NewSysAPIServiceClient,
	rpc.NewSysLogServiceClient,
	rpc.NewSysJobServiceClient,
	rpc.NewSysDeptServiceClient,
	rpc.NewRPCUserGrpc,
	rpc.NewUserServiceClient,
)

type Data struct {
	log        *log.Helper
	userClient userV1.UserClient
}

func NewData(c *conf.Bootstrap, logger log.Logger, userClient userV1.UserClient) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", fmt.Sprintf("%s/data", c.Name)))
	d := &Data{
		log:        l,
		userClient: userClient,
	}
	cleanup := func() {
		l.Info("closing the data resources")
	}
	return d, cleanup, nil
}
