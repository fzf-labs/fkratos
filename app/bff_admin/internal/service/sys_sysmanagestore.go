package service

import (
	"context"

	pb "fkratos/api/bff_admin/v1"
)

func (s *SysService) SysManageStore(ctx context.Context, req *pb.SysManageStoreReq) (*pb.SysManageStoreReply, error) {
	return s.sysUseCase.SysManageStore(ctx, req)
}
