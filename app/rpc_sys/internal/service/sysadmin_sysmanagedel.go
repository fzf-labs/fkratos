package service

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

// SysManageDel 管理员-删除
func (s *SysAdminService) SysManageDel(ctx context.Context, req *pb.SysManageDelReq) (*pb.SysManageDelReply, error) {
	return s.sysAdminUseCase.SysManageDel(ctx, req)
}
