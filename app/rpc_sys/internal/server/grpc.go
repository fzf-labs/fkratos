package server

import (
	pb "fkratos/api/rpc_sys/v1"
	"fkratos/app/rpc_sys/internal/service"

	conf "github.com/fzf-labs/fkratos-contrib/api/conf/v1"

	"github.com/fzf-labs/fkratos-contrib/bootstrap"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(
	c *conf.Bootstrap,
	logger log.Logger,
	authService *service.SysAuthService,
	adminService *service.SysAdminService,
	roleService *service.SysRoleService,
	permissionService *service.SysPermissionService,
	jobService *service.SysJobService,
	deptService *service.SysDeptService,
	apiService *service.SysAPIService,
	logService *service.SysLogService,
	dashboardService *service.DashboardService,
) *grpc.Server {
	// 创建grpc服务
	srv := bootstrap.NewGrpcServer(c, logger)

	// 注册服务
	pb.RegisterSysAuthServer(srv, authService)
	pb.RegisterSysAdminServer(srv, adminService)
	pb.RegisterSysRoleServer(srv, roleService)
	pb.RegisterSysPermissionServer(srv, permissionService)
	pb.RegisterSysJobServer(srv, jobService)
	pb.RegisterSysDeptServer(srv, deptService)
	pb.RegisterSysAPIServer(srv, apiService)
	pb.RegisterSysLogServer(srv, logService)
	pb.RegisterDashboardServer(srv, dashboardService)
	return srv
}
