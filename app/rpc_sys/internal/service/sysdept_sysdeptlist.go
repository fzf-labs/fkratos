package service

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

func (s *SysDeptService) SysDeptList(ctx context.Context, req *pb.SysDeptListReq) (*pb.SysDeptListReply, error) {
	return s.sysDeptUseCase.SysDeptList(ctx, req)
}
