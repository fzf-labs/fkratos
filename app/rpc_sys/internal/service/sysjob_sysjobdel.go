package service

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

func (s *SysJobService) SysJobDel(ctx context.Context, req *pb.SysJobDelReq) (*pb.SysJobDelReply, error) {
	return s.sysJobUseCase.SysJobDel(ctx, req)
}
