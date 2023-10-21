package service

import (
	"context"

	pb "fkratos/api/bff_admin/v1"
)

func (s *SysService) SysAPIStore(ctx context.Context, req *pb.SysAPIStoreReq) (*pb.SysAPIStoreReply, error) {
	return s.sysUseCase.SysAPIStore(ctx, req)
}
