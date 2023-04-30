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
	l := log.NewHelper(log.With(logger, "module", "auth/usecase/auth-service"))
	return &AuthUseCase{
		sysAdminRepo: sysAdminRepo,
		log:          l,
	}
}

type AuthUseCase struct {
	log          *log.Helper
	sysAdminRepo SysAdminRepo
}

func (a *AuthUseCase) SysAuthLoginCaptcha(ctx context.Context, req *v1.SysAuthLoginCaptchaReq) (*v1.SysAuthLoginCaptchaReply, error) {
	resp := new(v1.SysAuthLoginCaptchaReply)
	driver := base64Captcha.NewDriverDigit(80, 240, 6, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)
	captchaId, picPath, err := cp.Generate()
	if err != nil {
		return nil, err
	}
	resp.CaptchaId = captchaId
	resp.CaptchaImg = picPath
	return resp, nil
}

func (a *AuthUseCase) SysAuthLogin(ctx context.Context, req *v1.SysAuthLoginReq) (*v1.SysAuthLoginReply, error) {
	resp := new(v1.SysAuthLoginReply)
	//验证码
	verify := base64Captcha.DefaultMemStore.Verify(req.CaptchaId, req.VerifyCode, true)
	if !verify {
		return nil, v1.ErrorIncorrectVerificationCode("username:%s code:%s", req.Username, req.VerifyCode)
	}
	//用户校验
	sysAdmin, err := a.sysAdminRepo.SysAdminInfoByUsername(ctx, req.GetUsername())
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
	//颁发token
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

func (a *AuthUseCase) SysAuthLogout(ctx context.Context, req *v1.SysAuthLogoutReq) (*v1.SysAuthLogoutReply, error) {
	resp := new(v1.SysAuthLogoutReply)
	// todo jwtuid
	err := a.sysAdminRepo.ClearJwTToken(ctx, "")
	if err != nil {
		return nil, err
	}
	return resp, nil
}
