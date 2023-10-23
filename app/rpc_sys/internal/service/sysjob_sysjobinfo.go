package service

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

// SysJobInfo 岗位-单个岗位信息
func (s *SysJobService) SysJobInfo(ctx context.Context, req *pb.SysJobInfoReq) (*pb.SysJobInfoReply, error) {
	return s.sysJobUseCase.SysJobInfo(ctx, req)
}
