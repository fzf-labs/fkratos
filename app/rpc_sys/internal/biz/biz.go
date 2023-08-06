package biz

import (
	"context"
	"fkratos/api/paginator"
	v1 "fkratos/api/rpc_sys/v1"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_model"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_repo"

	"github.com/fzf-labs/fpkg/jwt"
	"github.com/fzf-labs/fpkg/page"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewAdminUseCase,
	NewApiUseCase,
	NewAuthUseCase,
	NewJobUseCase,
	NewDeptUseCase,
	NewLogUseCase,
	NewRoleUseCase,
	NewPermissionUseCase,
)

type SysAdminRepo interface {
	fkratos_sys_repo.ISysAdminRepo
	GenerateJwTToken(ctx context.Context, kv map[string]interface{}) (*jwt.Token, error)
	ClearJwTToken(ctx context.Context, jwtUId string) error
	SysManageListBySearch(ctx context.Context, req *paginator.PaginatorReq) ([]*fkratos_sys_model.SysAdmin, *page.Page, error)
	SysManageStore(ctx context.Context, req *v1.SysManageStoreReq) (*fkratos_sys_model.SysAdmin, error)
	GetAdminIdToNameByIds(ctx context.Context, ids []string) (map[string]string, error)
	SysAuthJwtTokenCheck(ctx context.Context, token string) (string, error)
}
type SysRoleRepo interface {
	fkratos_sys_repo.ISysRoleRepo
	GetRoleIdToNameByIds(ctx context.Context, ids []string) (map[string]string, error)
	SysRoleList(ctx context.Context) ([]*fkratos_sys_model.SysRole, error)
	SysRoleStore(ctx context.Context, req *v1.SysRoleStoreReq) (*fkratos_sys_model.SysRole, error)
}

type SysPermissionRepo interface {
	fkratos_sys_repo.ISysPermissionRepo
	SysPermissionList(ctx context.Context) ([]*fkratos_sys_model.SysPermission, error)
	SysPermissionStore(ctx context.Context, req *v1.SysPermissionStoreReq) (*fkratos_sys_model.SysPermission, error)
	SysPermissionUpdateStatus(ctx context.Context, id string, status int32) error
	SysPermissionByIdsAndStatus(ctx context.Context, ids []string, status int16) ([]*fkratos_sys_model.SysPermission, error)
	SysPermissionByStatus(ctx context.Context, status int16) ([]*fkratos_sys_model.SysPermission, error)
}

type SysJobRepo interface {
	fkratos_sys_repo.ISysJobRepo
	GetJobIdToNameByIds(ctx context.Context, ids []string) (map[string]string, error)
	SysJobListBySearch(ctx context.Context, req *paginator.PaginatorReq) ([]*fkratos_sys_model.SysJob, *page.Page, error)
	SysJobStore(ctx context.Context, req *v1.SysJobStoreReq) (*fkratos_sys_model.SysJob, error)
}

type SysDeptRepo interface {
	fkratos_sys_repo.ISysDeptRepo
	GetDeptIdToNameByIds(ctx context.Context, ids []string) (map[string]string, error)
	SysDeptList(ctx context.Context) ([]*v1.SysDeptInfo, error)
	SysDeptStore(ctx context.Context, req *v1.SysDeptStoreReq) (*fkratos_sys_model.SysDept, error)
}

type SysLogRepo interface {
	fkratos_sys_repo.ISysLogRepo
	SysLogListBySearch(ctx context.Context, req *paginator.PaginatorReq) ([]*fkratos_sys_model.SysLog, *page.Page, error)
	SysLogStore(ctx context.Context, req *v1.SysLogStoreReq) (*fkratos_sys_model.SysLog, error)
	SysLogStoreMQProducer(ctx context.Context, req *v1.SysLogStoreReq) error
}

type SysApiRepo interface {
	fkratos_sys_repo.ISysAPIRepo
	GetApiIdToNameByIds(ctx context.Context, ids []string) (map[string]string, error)
	SysApiStore(ctx context.Context, model *fkratos_sys_model.SysAPI) (*fkratos_sys_model.SysAPI, error)
}
