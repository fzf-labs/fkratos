package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewDashboardUseCase,
	NewSysAdminUseCase,
	NewSysAPIUseCase,
	NewSysAuthUseCase,
	NewSysDeptUseCase,
	NewSysJobUseCase,
	NewSysLogUseCase,
	NewSysPermissionUseCase,
	NewSysRoleUseCase,
)
