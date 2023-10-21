package service

import (
	"context"

	pb "fkratos/api/bff_admin/v1"
)

func (s *SysService) SysAPIList(ctx context.Context, req *pb.SysAPIListReq) (*pb.SysAPIListReply, error) {
	return s.sysUseCase.SysAPIList(ctx, req)
}
