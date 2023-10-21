package service

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

func (s *SysDeptService) SysDeptDel(ctx context.Context, req *pb.SysDeptDelReq) (*pb.SysDeptDelReply, error) {
	return s.sysDeptUseCase.SysDeptDel(ctx, req)
}
