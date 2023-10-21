package service

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

func (s *SysPermissionService) SysPermissionStatus(ctx context.Context, req *pb.SysPermissionStatusReq) (*pb.SysPermissionStatusResp, error) {
	return s.sysPermissionUseCase.SysPermissionStatus(ctx, req)
}
