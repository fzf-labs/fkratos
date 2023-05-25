package service

import "github.com/google/wire"

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	NewAuthService,
	NewAdminService,
	NewJobService,
	NewDeptService,
	NewLoginLogService,
	NewOperationLogService,
	NewRoleService,
	NewPermissionService,
	NewApiService,
	NewDashboardService,
)
