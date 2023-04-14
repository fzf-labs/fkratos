package bootstrap

import (
	"fkratos/internal/bootstrap/conf"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	kHttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
)

// NewHttpServer 创建Http服务端
func NewHttpServer(cfg *conf.Bootstrap, m ...middleware.Middleware) *kHttp.Server {
	var opts = []kHttp.ServerOption{
		kHttp.Filter(handlers.CORS(
			handlers.AllowedHeaders(cfg.Server.Http.Headers),
			handlers.AllowedMethods(cfg.Server.Http.Methods),
			handlers.AllowedOrigins(cfg.Server.Http.Origins),
		)),
	}
	var ms []middleware.Middleware
	ms = append(ms, recovery.Recovery())
	ms = append(ms, tracing.Server())
	ms = append(ms, m...)
	opts = append(opts, kHttp.Middleware(ms...))
	if cfg.Server.Http.Network != "" {
		opts = append(opts, kHttp.Network(cfg.Server.Http.Network))
	}
	if cfg.Server.Http.Addr != "" {
		opts = append(opts, kHttp.Address(cfg.Server.Http.Addr))
	}
	if cfg.Server.Http.Timeout != nil {
		opts = append(opts, kHttp.Timeout(cfg.Server.Http.Timeout.AsDuration()))
	}
	return kHttp.NewServer(opts...)
}
