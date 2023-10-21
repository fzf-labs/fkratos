package service

import (
	"context"

	pb "fkratos/api/bff_admin/v1"
)

func (s *SysService) SysDeptInfo(ctx context.Context, req *pb.SysDeptInfoReq) (*pb.SysDeptInfoReply, error) {
	return s.sysUseCase.SysDeptInfo(ctx, req)
}
