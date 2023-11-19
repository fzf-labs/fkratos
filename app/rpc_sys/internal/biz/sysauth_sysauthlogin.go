package biz

import (
	"context"
	"fkratos/internal/constant"
	"fkratos/internal/errorx"

	pb "fkratos/api/rpc_sys/v1"

	"github.com/fzf-labs/fpkg/crypt"
	"github.com/mojocn/base64Captcha"
)

// SysAuthLogin Auth-登录
func (s *SysAuthUseCase) SysAuthLogin(ctx context.Context, req *pb.SysAuthLoginReq) (*pb.SysAuthLoginReply, error) {
	resp := &pb.SysAuthLoginReply{}
	// 验证码
	verify := base64Captcha.DefaultMemStore.Verify(req.CaptchaId, req.VerifyCode, true)
	if !verify {
		return nil, errorx.AccountVerificationCodeErr.Err()
	}
	// 用户校验
	sysAdmin, err := s.sysAdminRepo.FindOneCacheByUsername(ctx, req.GetUsername())
	if err != nil {
		return nil, err
	}
	if sysAdmin == nil {
		return nil, errorx.AccountNotExist.Err()
	}
	if crypt.Compare(sysAdmin.Password, req.Password+sysAdmin.Salt) != nil {
		return nil, errorx.AccountWrongPassword.Err()
	}
	if sysAdmin.Status != constant.StatusEnable {
		return nil, errorx.AccountAbnormalStatus.Err()
	}
	// 颁发token
	kv := make(map[string]interface{})
	kv["uid"] = sysAdmin.ID
	token, err := s.sysAdminRepo.GenerateJwTToken(ctx, kv)
	if err != nil {
		return nil, err
	}
	resp.Token = token.Token
	resp.RefreshAt = token.RefreshAt
	resp.ExpiredAt = token.ExpiredAt
	return resp, nil
}
