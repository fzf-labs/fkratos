package service

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

// SysLogList 日志-列表
func (s *SysLogService) SysLogList(ctx context.Context, req *pb.SysLogListReq) (*pb.SysLogListResp, error) {
	return s.sysLogUseCase.SysLogList(ctx, req)
}
