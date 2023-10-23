package service

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

// SysJobList 岗位-列表
func (s *SysJobService) SysJobList(ctx context.Context, req *pb.SysJobListReq) (*pb.SysJobListReply, error) {
	return s.sysJobUseCase.SysJobList(ctx, req)
}
