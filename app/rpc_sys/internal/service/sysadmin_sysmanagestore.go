package service

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

func (s *SysAdminService) SysManageStore(ctx context.Context, req *pb.SysManageStoreReq) (*pb.SysManageStoreReply, error) {
	return s.sysAdminUseCase.SysManageStore(ctx, req)
}
