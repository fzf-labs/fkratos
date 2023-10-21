package biz

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

// SysPermissionDel 权限-删除
func (s *SysPermissionUseCase) SysPermissionDel(ctx context.Context, req *pb.SysPermissionDelReq) (*pb.SysPermissionDelResp, error) {
	resp := &pb.SysPermissionDelResp{}
	err := s.sysPermissionRepo.DeleteMultiCacheByIDS(ctx, req.GetIds())
	if err != nil {
		return nil, err
	}
	return resp, nil
}
