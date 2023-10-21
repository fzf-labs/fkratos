package service

import (
	"context"

	pb "fkratos/api/bff_admin/v1"
)

func (s *SysService) SysJobList(ctx context.Context, req *pb.SysJobListReq) (*pb.SysJobListReply, error) {
	return s.sysUseCase.SysJobList(ctx, req)
}
