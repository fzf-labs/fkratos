package service

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

// SysAdminInfoUpdate 管理员-个人信息更新
func (s *SysAdminService) SysAdminInfoUpdate(ctx context.Context, req *pb.SysAdminInfoUpdateReq) (*pb.SysAdminInfoUpdateReply, error) {
	return s.sysAdminUseCase.SysAdminInfoUpdate(ctx, req)
}
