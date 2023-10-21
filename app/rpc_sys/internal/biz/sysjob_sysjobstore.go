package biz

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

// SysJobStore 岗位-保存
func (s *SysJobUseCase) SysJobStore(ctx context.Context, req *pb.SysJobStoreReq) (*pb.SysJobStoreReply, error) {
	resp := &pb.SysJobStoreReply{}
	_, err := s.sysJobRepo.SysJobStore(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
