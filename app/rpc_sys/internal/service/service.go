package service

import "github.com/google/wire"

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	NewAuthService,
	NewAdminService,
	NewJobService,
	NewDeptService,
	NewLogService,
	NewRoleService,
	NewPermissionService,
	NewAPIService,
	NewDashboardService,
)
