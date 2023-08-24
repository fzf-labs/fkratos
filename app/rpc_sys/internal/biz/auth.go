package biz

import (
	"context"
	v1 "fkratos/api/rpc_sys/v1"
	"fkratos/internal/constant"
	"fkratos/internal/errorx"

	"github.com/fzf-labs/fpkg/crypt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/mojocn/base64Captcha"
)

func NewAuthUseCase(logger log.Logger, sysAdminRepo SysAdminRepo) *AuthUseCase {
	l := log.NewHelper(log.With(logger, "module", "rpc_user/biz/auth"))
	return &AuthUseCase{
		sysAdminRepo: sysAdminRepo,
		log:          l,
	}
}

type AuthUseCase struct {
	log          *log.Helper
	sysAdminRepo SysAdminRepo
}

// SysAuthLoginCaptcha 验证码
func (a *AuthUseCase) SysAuthLoginCaptcha(_ context.Context, _ *v1.SysAuthLoginCaptchaReq) (*v1.SysAuthLoginCaptchaReply, error) {
	resp := new(v1.SysAuthLoginCaptchaReply)
	driver := base64Captcha.NewDriverDigit(80, 240, 6, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)
	captchaID, picPath, err := cp.Generate()
	if err != nil {
		return nil, err
	}
	resp.CaptchaId = captchaID
	resp.CaptchaImg = picPath
	return resp, nil
}

// SysAuthLogin 登录
func (a *AuthUseCase) SysAuthLogin(ctx context.Context, req *v1.SysAuthLoginReq) (*v1.SysAuthLoginReply, error) {
	resp := new(v1.SysAuthLoginReply)
	// 验证码
	verify := base64Captcha.DefaultMemStore.Verify(req.CaptchaId, req.VerifyCode, true)
	if !verify {
		return nil, errorx.AccountVerificationCodeErr
	}
	// 用户校验
	sysAdmin, err := a.sysAdminRepo.FindOneCacheByUsername(ctx, req.GetUsername())
	if err != nil {
		return nil, err
	}
	if sysAdmin == nil {
		return nil, errorx.AccountNotExist
	}
	if crypt.Compare(sysAdmin.Password, req.Password+sysAdmin.Salt) != nil {
		return nil, errorx.AccountWrongPassword
	}
	if sysAdmin.Status != constant.StatusEnable {
		return nil, errorx.AccountAbnormalStatus
	}
	// 颁发token
	kv := make(map[string]interface{})
	kv["uid"] = sysAdmin.ID
	token, err := a.sysAdminRepo.GenerateJwTToken(ctx, kv)
	if err != nil {
		return nil, err
	}
	resp.Token = token.Token
	resp.RefreshAt = token.RefreshAt
	resp.ExpiredAt = token.ExpiredAt
	return resp, nil
}

// SysAuthLogout 登出
func (a *AuthUseCase) SysAuthLogout(ctx context.Context, req *v1.SysAuthLogoutReq) (*v1.SysAuthLogoutReply, error) {
	resp := new(v1.SysAuthLogoutReply)
	err := a.sysAdminRepo.ClearJwTToken(ctx, req.GetAdminId())
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SysAuthJwtTokenCheck token校验
func (a *AuthUseCase) SysAuthJwtTokenCheck(ctx context.Context, req *v1.SysAuthJwtTokenCheckReq) (*v1.SysAuthJwtTokenCheckReply, error) {
	resp := new(v1.SysAuthJwtTokenCheckReply)
	adminID, err := a.sysAdminRepo.SysAuthJwtTokenCheck(ctx, req.GetToken())
	if err != nil {
		return nil, err
	}
	resp.AdminId = adminID
	return resp, nil
}
