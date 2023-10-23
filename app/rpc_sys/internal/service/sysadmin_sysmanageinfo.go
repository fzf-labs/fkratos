package service

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

// SysManageInfo 管理员-信息
func (s *SysAdminService) SysManageInfo(ctx context.Context, req *pb.SysManageInfoReq) (*pb.SysManageInfoReply, error) {
	return s.sysAdminUseCase.SysManageInfo(ctx, req)
}
