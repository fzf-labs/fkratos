package service

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

// SysAuthLoginCaptcha Auth-验证码
func (s *SysAuthService) SysAuthLoginCaptcha(ctx context.Context, req *pb.SysAuthLoginCaptchaReq) (*pb.SysAuthLoginCaptchaReply, error) {
	return s.sysAuthUseCase.SysAuthLoginCaptcha(ctx, req)
}
