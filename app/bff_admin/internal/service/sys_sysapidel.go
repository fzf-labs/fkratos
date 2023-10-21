package service

import (
	"context"

	pb "fkratos/api/bff_admin/v1"
)

func (s *SysService) SysAPIDel(ctx context.Context, req *pb.SysAPIDelReq) (*pb.SysAPIDelReply, error) {
	return s.sysUseCase.SysAPIDel(ctx, req)
}
