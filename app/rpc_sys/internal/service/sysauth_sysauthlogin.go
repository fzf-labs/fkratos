package service

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

// SysAuthLogin Auth-登录
func (s *SysAuthService) SysAuthLogin(ctx context.Context, req *pb.SysAuthLoginReq) (*pb.SysAuthLoginReply, error) {
	return s.sysAuthUseCase.SysAuthLogin(ctx, req)
}
