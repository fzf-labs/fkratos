package service

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

// SysPermissionInfo 权限-单个权限信息
func (s *SysPermissionService) SysPermissionInfo(ctx context.Context, req *pb.SysPermissionInfoReq) (*pb.SysPermissionInfoResp, error) {
	return s.sysPermissionUseCase.SysPermissionInfo(ctx, req)
}
