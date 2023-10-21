package biz

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

// SysDeptList 部门-列表
func (s *SysDeptUseCase) SysDeptList(ctx context.Context, _ *pb.SysDeptListReq) (*pb.SysDeptListReply, error) {
	resp := &pb.SysDeptListReply{}
	list, err := s.sysDeptRepo.SysDeptList(ctx)
	if err != nil {
		return nil, err
	}
	resp.List = list
	return resp, nil
}
