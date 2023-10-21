package service

import (
	"context"

	pb "fkratos/api/bff_admin/v1"
)

func (s *SysService) SysManageList(ctx context.Context, req *pb.SysManageListReq) (*pb.SysManageListReply, error) {
	return s.sysUseCase.SysManageList(ctx, req)
}
