package service

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

// SysPermissionStatus 权限-修改状态
func (s *SysPermissionService) SysPermissionStatus(ctx context.Context, req *pb.SysPermissionStatusReq) (*pb.SysPermissionStatusResp, error) {
	return s.sysPermissionUseCase.SysPermissionStatus(ctx, req)
}
