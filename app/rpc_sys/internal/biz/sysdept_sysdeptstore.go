package biz

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

// SysDeptStore 部门-保存
func (s *SysDeptUseCase) SysDeptStore(ctx context.Context, req *pb.SysDeptStoreReq) (*pb.SysDeptStoreReply, error) {
	resp := &pb.SysDeptStoreReply{}
	_, err := s.sysDeptRepo.SysDeptStore(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
