package service

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

func (s *SysRoleService) SysRoleList(ctx context.Context, req *pb.SysRoleListReq) (*pb.SysRoleListResp, error) {
	return s.sysRoleUseCase.SysRoleList(ctx, req)
}
