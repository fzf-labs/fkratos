package service

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

func (s *SysAuthService) SysAuthLogout(ctx context.Context, req *pb.SysAuthLogoutReq) (*pb.SysAuthLogoutReply, error) {
	return s.sysAuthUseCase.SysAuthLogout(ctx, req)
}
