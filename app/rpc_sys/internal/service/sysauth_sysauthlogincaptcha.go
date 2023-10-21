package service

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

func (s *SysAuthService) SysAuthLoginCaptcha(ctx context.Context, req *pb.SysAuthLoginCaptchaReq) (*pb.SysAuthLoginCaptchaReply, error) {
	return s.sysAuthUseCase.SysAuthLoginCaptcha(ctx, req)
}
