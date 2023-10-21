package biz

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

// SysRoleDel 角色-删除
func (s *SysRoleUseCase) SysRoleDel(ctx context.Context, req *pb.SysRoleDelReq) (*pb.SysRoleDelResp, error) {
	resp := &pb.SysRoleDelResp{}
	err := s.sysRoleRepo.DeleteMultiCacheByIDS(ctx, req.GetIds())
	if err != nil {
		return nil, err
	}
	return resp, nil
}
