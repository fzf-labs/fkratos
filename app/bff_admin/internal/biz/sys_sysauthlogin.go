package biz

import (
	"context"
	sysV1 "fkratos/api/rpc_sys/v1"
	"fkratos/internal/errorx"

	pb "fkratos/api/bff_admin/v1"

	"github.com/jinzhu/copier"
)

// SysAuthLogin Auth-登录
func (s *SysUseCase) SysAuthLogin(ctx context.Context, req *pb.SysAuthLoginReq) (*pb.SysAuthLoginReply, error) {
	resp := new(pb.SysAuthLoginReply)
	result, err := s.sysAuthClient.SysAuthLogin(ctx, &sysV1.SysAuthLoginReq{
		CaptchaId:  req.GetCaptchaId(),
		VerifyCode: req.GetVerifyCode(),
		Username:   req.GetUsername(),
		Password:   req.GetPassword(),
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(resp, result)
	if err != nil {
		return nil, errorx.DataFormattingError.WithError(err).Err()
	}
	return resp, nil
}
