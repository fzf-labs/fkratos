package biz

import (
	"context"
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
)

type SysAdminRepo interface {
	SysAdminInfoByUsername(ctx context.Context, username string) (*rpc_sys_model.SysAdmin, error)
	SysAdminInfoByAdminId(ctx context.Context, adminId string) (*rpc_sys_model.SysAdmin, error)
	SysAdminInfoCacheByAdminId(ctx context.Context, adminId string) (*v1.SysAdminInfo, error)
	SysAdminInfoUpdate(ctx context.Context, req *v1.SysAdminInfoUpdateReq) (*v1.SysAdminInfoUpdateReply, error)
	SysAdminDel(ctx context.Context, ids []string) error
	GenerateJwTToken(ctx context.Context, kv map[string]interface{}) (*jwt.Token, error)
	ClearJwTToken(ctx context.Context, jwtUId string) error
	SysManageListBySearch(ctx context.Context, req *v1.SysManageListReq) ([]*rpc_sys_model.SysAdmin, *page.Page, error)
	SysManageStore(ctx context.Context, req *v1.SysManageStoreReq) (*rpc_sys_model.SysAdmin, error)
}

type SysRoleRepo interface {
	GetRoleIdToNameByIds(ctx context.Context, ids []string) (map[string]string, error)
}

type SysJobRepo interface {
	GetJobIdToNameByIds(ctx context.Context, ids []string) (map[string]string, error)
}

type SysDeptRepo interface {
	GetDeptIdToNameByIds(ctx context.Context, ids []string) (map[string]string, error)
}
