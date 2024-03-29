package server

import (
	bffAdminV1 "fkratos/api/bff_admin/v1"
	sysV1 "fkratos/api/rpc_sys/v1"
	"fkratos/app/bff_admin/internal/middleware/auth"
	"fkratos/app/bff_admin/internal/service"

	conf "github.com/fzf-labs/fkratos-contrib/api/conf/v1"

	"github.com/fzf-labs/fkratos-contrib/bootstrap"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Bootstrap, logger log.Logger, authClient sysV1.SysAuthClient, sysService *service.SysService) *http.Server {
	srv := bootstrap.NewHTTPServer(c, logger, auth.SelectorMiddleware(authClient))
	bffAdminV1.RegisterSysHTTPServer(srv, sysService)
	return srv
}
