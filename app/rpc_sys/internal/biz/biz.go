package biz

import (
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewAdminUseCase,
	NewAPIUseCase,
	NewAuthUseCase,
	NewJobUseCase,
	NewDeptUseCase,
	NewLogUseCase,
	NewRoleUseCase,
	NewPermissionUseCase,
)
