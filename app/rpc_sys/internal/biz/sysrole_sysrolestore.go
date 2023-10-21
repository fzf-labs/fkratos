package biz

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

// SysRoleStore 角色-保存
func (s *SysRoleUseCase) SysRoleStore(ctx context.Context, req *pb.SysRoleStoreReq) (*pb.SysRoleStoreResp, error) {
	resp := &pb.SysRoleStoreResp{}
	_, err := s.sysRoleRepo.SysRoleStore(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
