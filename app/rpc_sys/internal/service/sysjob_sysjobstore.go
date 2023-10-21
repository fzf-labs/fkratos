package service

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

func (s *SysJobService) SysJobStore(ctx context.Context, req *pb.SysJobStoreReq) (*pb.SysJobStoreReply, error) {
	return s.sysJobUseCase.SysJobStore(ctx, req)
}
