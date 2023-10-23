package service

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

// SysManageList 管理员-列表
func (s *SysAdminService) SysManageList(ctx context.Context, req *pb.SysManageListReq) (*pb.SysManageListReply, error) {
	return s.sysAdminUseCase.SysManageList(ctx, req)
}
