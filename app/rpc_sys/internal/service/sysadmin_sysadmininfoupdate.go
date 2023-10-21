package service

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

func (s *SysAdminService) SysAdminInfoUpdate(ctx context.Context, req *pb.SysAdminInfoUpdateReq) (*pb.SysAdminInfoUpdateReply, error) {
	return s.sysAdminUseCase.SysAdminInfoUpdate(ctx, req)
}
