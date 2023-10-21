package service

import (
	"context"

	pb "fkratos/api/bff_admin/v1"
)

func (s *SysService) SysJobInfo(ctx context.Context, req *pb.SysJobInfoReq) (*pb.SysJobInfoReply, error) {
	return s.sysUseCase.SysJobInfo(ctx, req)
}
