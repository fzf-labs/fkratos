package service

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

// SysLogInfo 日志-信息
func (s *SysLogService) SysLogInfo(ctx context.Context, req *pb.SysLogInfoReq) (*pb.SysLogInfoResp, error) {
	return s.sysLogUseCase.SysLogInfo(ctx, req)
}
