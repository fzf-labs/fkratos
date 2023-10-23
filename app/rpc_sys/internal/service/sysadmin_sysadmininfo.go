package service

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

// SysAdminInfo 管理员-个人信息
func (s *SysAdminService) SysAdminInfo(ctx context.Context, req *pb.SysAdminInfoReq) (*pb.SysAdminInfoReply, error) {
	return s.sysAdminUseCase.SysAdminInfo(ctx, req)
}
