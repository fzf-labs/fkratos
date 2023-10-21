package biz

import (
	sysV1 "fkratos/api/rpc_sys/v1"

	"github.com/go-kratos/kratos/v2/log"
)

func NewSysUseCase(
	logger log.Logger,
	sysAuthClient sysV1.SysAuthClient,
	sysAdminClient sysV1.SysAdminClient,
	sysRoleClient sysV1.SysRoleClient,
	sysPermissionClient sysV1.SysPermissionClient,
	sysJobClient sysV1.SysJobClient,
	sysDeptClient sysV1.SysDeptClient,
	sysAPIClient sysV1.SysAPIClient,
	sysLogClient sysV1.SysLogClient,
	sysDashboardClient sysV1.DashboardClient,
) *SysUseCase {
	l := log.NewHelper(log.With(logger, "module", "biz/sys"))
	return &SysUseCase{
		log:                 l,
		sysAuthClient:       sysAuthClient,
		sysAdminClient:      sysAdminClient,
		sysRoleClient:       sysRoleClient,
		sysPermissionClient: sysPermissionClient,
		sysJobClient:        sysJobClient,
		sysDeptClient:       sysDeptClient,
		sysAPIClient:        sysAPIClient,
		sysLogClient:        sysLogClient,
		sysDashboardClient:  sysDashboardClient,
	}
}

type SysUseCase struct {
	log                 *log.Helper
	sysAuthClient       sysV1.SysAuthClient
	sysAdminClient      sysV1.SysAdminClient
	sysRoleClient       sysV1.SysRoleClient
	sysPermissionClient sysV1.SysPermissionClient
	sysJobClient        sysV1.SysJobClient
	sysDeptClient       sysV1.SysDeptClient
	sysAPIClient        sysV1.SysAPIClient
	sysLogClient        sysV1.SysLogClient
	sysDashboardClient  sysV1.DashboardClient
}
