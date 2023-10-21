package service

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

func (s *SysRoleService) SysRoleStore(ctx context.Context, req *pb.SysRoleStoreReq) (*pb.SysRoleStoreResp, error) {
	return s.sysRoleUseCase.SysRoleStore(ctx, req)
}
