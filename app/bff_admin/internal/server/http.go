package server

import (
	"fkratos/api/bff_admin/v1"
	"fkratos/app/bff_admin/internal/service"
	"fkratos/internal/bootstrap"
	"fkratos/internal/bootstrap/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Bootstrap, logger log.Logger, adminService *service.AdminService) *http.Server {
	srv := bootstrap.NewHttpServer(c, logger)
	v1.RegisterAdminHTTPServer(srv, adminService)
	return srv
}
