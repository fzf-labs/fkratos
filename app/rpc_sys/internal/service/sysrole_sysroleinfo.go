package service

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

// SysRoleInfo 角色-信息
func (s *SysRoleService) SysRoleInfo(ctx context.Context, req *pb.SysRoleInfoReq) (*pb.SysRoleInfoResp, error) {
	return s.sysRoleUseCase.SysRoleInfo(ctx, req)
}
