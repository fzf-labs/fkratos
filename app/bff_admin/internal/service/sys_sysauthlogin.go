package service

import (
	"context"

	pb "fkratos/api/bff_admin/v1"
)

func (s *SysService) SysAuthLogin(ctx context.Context, req *pb.SysAuthLoginReq) (*pb.SysAuthLoginReply, error) {
	return s.sysUseCase.SysAuthLogin(ctx, req)
}
