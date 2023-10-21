package biz

import (
	"context"
	sysV1 "fkratos/api/rpc_sys/v1"
	"fkratos/internal/dto"

	pb "fkratos/api/bff_admin/v1"
)

// SysAuthLoginCaptcha Auth-验证码
func (s *SysUseCase) SysAuthLoginCaptcha(ctx context.Context, _ *pb.SysAuthLoginCaptchaReq) (*pb.SysAuthLoginCaptchaReply, error) {
	resp := &pb.SysAuthLoginCaptchaReply{}
	captcha, err := s.sysAuthClient.SysAuthLoginCaptcha(ctx, &sysV1.SysAuthLoginCaptchaReq{})
	if err != nil {
		return nil, err
	}
	err = dto.Copy(resp, captcha)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
