package service

import "github.com/google/wire"

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	NewDashboardService,
	NewSysAdminService,
	NewSysAPIService,
	NewSysAuthService,
	NewSysDeptService,
	NewSysJobService,
	NewSysLogService,
	NewSysPermissionService,
	NewSysRoleService,
)
