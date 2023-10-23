package service

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

// SysLogStore 日志-保存
func (s *SysLogService) SysLogStore(ctx context.Context, req *pb.SysLogStoreReq) (*pb.SysLogStoreResp, error) {
	return s.sysLogUseCase.SysLogStore(ctx, req)
}
