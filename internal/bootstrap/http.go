package bootstrap

import (
	"fkratos/internal/bootstrap/conf"
	"fkratos/internal/errorx"
	"fkratos/internal/middleware/ctx"
	"fkratos/internal/middleware/validate"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// NewHTTPServer 创建Http服务端
func NewHTTPServer(cfg *conf.Bootstrap, logger log.Logger, m ...middleware.Middleware) *http.Server {
	var opts = []http.ServerOption{
		http.Filter(handlers.CORS(
			handlers.AllowedHeaders(cfg.Server.Http.Headers),
			handlers.AllowedMethods(cfg.Server.Http.Methods),
			handlers.AllowedOrigins(cfg.Server.Http.Origins),
		)),
	}
	var ms []middleware.Middleware
	ms = append(ms,
		tracing.Server(),
		logging.Server(logger),
		recovery.Recovery(),
		metadata.Server(),
		validate.Validator(),
		ctx.Canceled(),
		Metrics(),
	)
	ms = append(ms, m...)
	opts = append(opts, http.Middleware(ms...))
	if cfg.Server.Http.Network != "" {
		opts = append(opts, http.Network(cfg.Server.Http.Network))
	}
	if cfg.Server.Http.Addr != "" {
		opts = append(opts, http.Address(cfg.Server.Http.Addr))
	}
	if cfg.Server.Http.Timeout != nil {
		opts = append(opts, http.Timeout(cfg.Server.Http.Timeout.AsDuration()))
	}
	opts = append(opts, http.ErrorEncoder(errorx.ErrorEncoder))
	srv := http.NewServer(opts...)
	srv.Handle("/metrics", promhttp.Handler())
	return srv
}
