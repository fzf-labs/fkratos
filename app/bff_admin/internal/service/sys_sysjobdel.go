package service

import (
	"context"

	pb "fkratos/api/bff_admin/v1"
)

func (s *SysService) SysJobDel(ctx context.Context, req *pb.SysJobDelReq) (*pb.SysJobDelReply, error) {
	return s.sysUseCase.SysJobDel(ctx, req)
}
