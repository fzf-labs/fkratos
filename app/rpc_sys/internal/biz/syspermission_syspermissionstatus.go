package biz

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

// SysPermissionStatus 权限-修改状态
func (s *SysPermissionUseCase) SysPermissionStatus(ctx context.Context, req *pb.SysPermissionStatusReq) (*pb.SysPermissionStatusResp, error) {
	resp := &pb.SysPermissionStatusResp{}
	err := s.sysPermissionRepo.SysPermissionUpdateStatus(ctx, req.GetId(), req.GetStatus())
	if err != nil {
		return nil, err
	}
	return resp, nil
}
