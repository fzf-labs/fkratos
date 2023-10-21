package biz

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

// SysPermissionInfo 权限-单个权限信息
func (s *SysPermissionUseCase) SysPermissionInfo(ctx context.Context, req *pb.SysPermissionInfoReq) (*pb.SysPermissionInfoResp, error) {
	resp := &pb.SysPermissionInfoResp{}
	_, err := s.sysPermissionRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return resp, nil
}
