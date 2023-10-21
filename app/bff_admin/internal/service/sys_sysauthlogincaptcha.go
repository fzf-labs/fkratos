package service

import (
	"context"

	pb "fkratos/api/bff_admin/v1"
)

func (s *SysService) SysAuthLoginCaptcha(ctx context.Context, req *pb.SysAuthLoginCaptchaReq) (*pb.SysAuthLoginCaptchaReply, error) {
	return s.sysUseCase.SysAuthLoginCaptcha(ctx, req)
}
