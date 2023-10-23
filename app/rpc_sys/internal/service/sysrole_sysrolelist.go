package service

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

// SysRoleList 角色-列表
func (s *SysRoleService) SysRoleList(ctx context.Context, req *pb.SysRoleListReq) (*pb.SysRoleListResp, error) {
	return s.sysRoleUseCase.SysRoleList(ctx, req)
}
