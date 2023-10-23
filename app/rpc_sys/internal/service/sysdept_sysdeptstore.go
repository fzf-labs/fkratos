package service

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

// SysDeptStore 部门-保存
func (s *SysDeptService) SysDeptStore(ctx context.Context, req *pb.SysDeptStoreReq) (*pb.SysDeptStoreReply, error) {
	return s.sysDeptUseCase.SysDeptStore(ctx, req)
}
