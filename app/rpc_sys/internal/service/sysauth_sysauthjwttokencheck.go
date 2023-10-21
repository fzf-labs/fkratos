package service

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

func (s *SysAuthService) SysAuthJwtTokenCheck(ctx context.Context, req *pb.SysAuthJwtTokenCheckReq) (*pb.SysAuthJwtTokenCheckReply, error) {
	return s.sysAuthUseCase.SysAuthJwtTokenCheck(ctx, req)
}
