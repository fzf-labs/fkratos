package service

import (
	"context"

	pb "fkratos/api/bff_admin/v1"
)

func (s *SysService) SysLogList(ctx context.Context, req *pb.SysLogListReq) (*pb.SysLogListResp, error) {
	return s.sysUseCase.SysLogList(ctx, req)
}
