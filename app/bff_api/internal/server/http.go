package server

import (
	bffAPIV1 "fkratos/api/bff_api/v1"
	"fkratos/app/bff_api/internal/service"

	conf "github.com/fzf-labs/fkratos-contrib/api/conf/v1"

	"github.com/fzf-labs/fkratos-contrib/bootstrap"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Bootstrap, logger log.Logger, userService *service.UserService) *http.Server {
	srv := bootstrap.NewHTTPServer(c, logger)
	bffAPIV1.RegisterUserHTTPServer(srv, userService)
	return srv
}
