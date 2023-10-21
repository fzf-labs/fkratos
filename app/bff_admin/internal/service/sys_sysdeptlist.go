package service

import (
	"context"

	pb "fkratos/api/bff_admin/v1"
)

func (s *SysService) SysDeptList(ctx context.Context, req *pb.SysDeptListReq) (*pb.SysDeptListReply, error) {
	return s.sysUseCase.SysDeptList(ctx, req)
}
