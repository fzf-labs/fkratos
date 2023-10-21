package biz

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

// SysDeptDel 部门-删除
func (s *SysDeptUseCase) SysDeptDel(ctx context.Context, req *pb.SysDeptDelReq) (*pb.SysDeptDelReply, error) {
	resp := &pb.SysDeptDelReply{}
	err := s.sysDeptRepo.DeleteMultiCacheByIDS(ctx, req.GetIds())
	if err != nil {
		return nil, err
	}
	return resp, nil
}
