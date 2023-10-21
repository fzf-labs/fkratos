package biz

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

// SysPermissionStore 权限-保存
func (s *SysPermissionUseCase) SysPermissionStore(ctx context.Context, req *pb.SysPermissionStoreReq) (*pb.SysPermissionStoreResp, error) {
	resp := &pb.SysPermissionStoreResp{}
	_, err := s.sysPermissionRepo.SysPermissionStore(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
