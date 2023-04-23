package server

import (
	"fkratos/api/bff_api/v1"
	"fkratos/app/bff_api/internal/service"
	"fkratos/internal/bootstrap"
	"fkratos/internal/bootstrap/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Bootstrap, logger log.Logger, apiService *service.ApiService) *http.Server {
	srv := bootstrap.NewHttpServer(c, logging.Server(logger))
	v1.RegisterApiHTTPServer(srv, apiService)
	return srv
}
