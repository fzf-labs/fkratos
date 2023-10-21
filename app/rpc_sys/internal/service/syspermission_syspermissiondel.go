package service

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

func (s *SysPermissionService) SysPermissionDel(ctx context.Context, req *pb.SysPermissionDelReq) (*pb.SysPermissionDelResp, error) {
	return s.sysPermissionUseCase.SysPermissionDel(ctx, req)
}
