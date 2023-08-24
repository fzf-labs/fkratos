package server

import (
	v1 "fkratos/api/rpc_sys/v1"
	"fkratos/app/rpc_sys/internal/service"
	"fkratos/internal/bootstrap"
	"fkratos/internal/bootstrap/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(
	c *conf.Bootstrap,
	logger log.Logger,
	authService *service.AuthService,
	adminService *service.AdminService,
	roleService *service.RoleService,
	permissionService *service.PermissionService,
	jobService *service.JobService,
	deptService *service.DeptService,
	apiService *service.APIService,
	logService *service.LogService,
	dashboardService *service.DashboardService,
) *grpc.Server {
	// 创建grpc服务
	srv := bootstrap.NewGrpcServer(c, logger)

	// 注册服务
	v1.RegisterAuthServer(srv, authService)
	v1.RegisterAdminServer(srv, adminService)
	v1.RegisterRoleServer(srv, roleService)
	v1.RegisterPermissionServer(srv, permissionService)
	v1.RegisterJobServer(srv, jobService)
	v1.RegisterDeptServer(srv, deptService)
	v1.RegisterAPIServer(srv, apiService)
	v1.RegisterLogServer(srv, logService)
	v1.RegisterDashboardServer(srv, dashboardService)
	return srv
}
