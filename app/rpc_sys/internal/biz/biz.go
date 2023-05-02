package biz

import (
	"context"
	"fkratos/api/common"
	v1 "fkratos/api/rpc_sys/v1"
	"fkratos/app/rpc_sys/internal/data/gorm/rpc_sys_model"

	"github.com/fzf-labs/fpkg/jwt"
	"github.com/fzf-labs/fpkg/page"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewAuthUseCase,
	NewAdminUseCase,
	NewJobUseCase,
	NewDeptUseCase,
	NewLogUseCase,
	NewRoleUseCase,
)

type SysAdminRepo interface {
	SysAdminInfoByUsername(ctx context.Context, username string) (*rpc_sys_model.SysAdmin, error)
	SysAdminInfoByAdminId(ctx context.Context, adminId string) (*rpc_sys_model.SysAdmin, error)
	SysAdminInfoCacheByAdminId(ctx context.Context, adminId string) (*v1.SysAdminInfo, error)
	SysAdminInfoUpdate(ctx context.Context, req *v1.SysAdminInfoUpdateReq) (*v1.SysAdminInfoUpdateReply, error)
	SysAdminDel(ctx context.Context, ids []string) error
	GenerateJwTToken(ctx context.Context, kv map[string]interface{}) (*jwt.Token, error)
	ClearJwTToken(ctx context.Context, jwtUId string) error
	SysManageListBySearch(ctx context.Context, req *common.SearchListReq) ([]*rpc_sys_model.SysAdmin, *page.Page, error)
	SysManageStore(ctx context.Context, req *v1.SysManageStoreReq) (*rpc_sys_model.SysAdmin, error)
	GetAdminIdToNameByIds(ctx context.Context, ids []string) (map[string]string, error)
}

type SysRoleRepo interface {
	GetRoleIdToNameByIds(ctx context.Context, ids []string) (map[string]string, error)
	SysRoleList(ctx context.Context) ([]*rpc_sys_model.SysRole, error)
	SysRoleInfoById(ctx context.Context, id string) (*rpc_sys_model.SysRole, error)
	SysRoleDelByIds(ctx context.Context, ids []string) error
	SysRoleStore(ctx context.Context, req *v1.SysRoleStoreReq) (*rpc_sys_model.SysRole, error)
}

type SysPermissionRepo interface {
	SysPermissionList(ctx context.Context) ([]*rpc_sys_model.SysPermission, error)
	SysPermissionInfoById(ctx context.Context, id string) (*rpc_sys_model.SysPermission, error)
	SysPermissionDelByIds(ctx context.Context, ids []string) error
	SysPermissionStore(ctx context.Context, req *v1.SysPermissionStoreReq) (*rpc_sys_model.SysPermission, error)
	SysPermissionUpdateStatus(ctx context.Context, id string, status int32) error
}

type SysJobRepo interface {
	GetJobIdToNameByIds(ctx context.Context, ids []string) (map[string]string, error)
	SysJobListBySearch(ctx context.Context, req *common.SearchListReq) ([]*rpc_sys_model.SysJob, *page.Page, error)
	SysJobInfoById(ctx context.Context, id string) (*v1.SysJobInfo, error)
	SysJobDelByIds(ctx context.Context, ids []string) error
	SysJobStore(ctx context.Context, req *v1.SysJobStoreReq) (*rpc_sys_model.SysJob, error)
}

type SysDeptRepo interface {
	GetDeptIdToNameByIds(ctx context.Context, ids []string) (map[string]string, error)
	SysDeptList(ctx context.Context) ([]*v1.SysDeptInfo, error)
	SysDeptInfoById(ctx context.Context, id string) (*v1.SysDeptInfo, error)
	SysDeptDelByIds(ctx context.Context, ids []string) error
	SysDeptStore(ctx context.Context, req *v1.SysDeptStoreReq) (*rpc_sys_model.SysDept, error)
}

type SysLogRepo interface {
	SysLogListBySearch(ctx context.Context, req *common.SearchListReq) ([]*rpc_sys_model.SysLog, *page.Page, error)
	SysLogInfoById(ctx context.Context, id string) (*rpc_sys_model.SysLog, error)
}

type SysApiRepo interface {
	GetApiIdToNameByIds(ctx context.Context, ids []string) (map[string]string, error)
	SysDeptList(ctx context.Context, permissionId string) ([]*rpc_sys_model.SysAPI, error)
}
