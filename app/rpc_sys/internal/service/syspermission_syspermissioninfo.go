package service

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

func (s *SysPermissionService) SysPermissionInfo(ctx context.Context, req *pb.SysPermissionInfoReq) (*pb.SysPermissionInfoResp, error) {
	return s.sysPermissionUseCase.SysPermissionInfo(ctx, req)
}
