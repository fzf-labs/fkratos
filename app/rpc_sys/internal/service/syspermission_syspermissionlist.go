package service

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

func (s *SysPermissionService) SysPermissionList(ctx context.Context, req *pb.SysPermissionListReq) (*pb.SysPermissionListResp, error) {
	return s.sysPermissionUseCase.SysPermissionList(ctx, req)
}
