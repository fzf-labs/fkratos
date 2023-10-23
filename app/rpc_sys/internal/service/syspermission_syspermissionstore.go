package service

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

// SysPermissionStore 权限-保存
func (s *SysPermissionService) SysPermissionStore(ctx context.Context, req *pb.SysPermissionStoreReq) (*pb.SysPermissionStoreResp, error) {
	return s.sysPermissionUseCase.SysPermissionStore(ctx, req)
}
