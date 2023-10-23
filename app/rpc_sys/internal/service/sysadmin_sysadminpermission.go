package service

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

// SysAdminPermission 管理员-查询权限
func (s *SysAdminService) SysAdminPermission(ctx context.Context, req *pb.SysAdminPermissionReq) (*pb.SysAdminPermissionReply, error) {
	return s.sysAdminUseCase.SysAdminPermission(ctx, req)
}
