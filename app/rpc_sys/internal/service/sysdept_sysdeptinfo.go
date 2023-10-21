package service

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

func (s *SysDeptService) SysDeptInfo(ctx context.Context, req *pb.SysDeptInfoReq) (*pb.SysDeptInfoReply, error) {
	return s.sysDeptUseCase.SysDeptInfo(ctx, req)
}
