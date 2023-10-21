package service

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

func (s *SysRoleService) SysRoleDel(ctx context.Context, req *pb.SysRoleDelReq) (*pb.SysRoleDelResp, error) {
	return s.sysRoleUseCase.SysRoleDel(ctx, req)
}
