package biz

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"

	"github.com/mojocn/base64Captcha"
)

// SysAuthLoginCaptcha Auth-验证码
func (s *SysAuthUseCase) SysAuthLoginCaptcha(_ context.Context, _ *pb.SysAuthLoginCaptchaReq) (*pb.SysAuthLoginCaptchaReply, error) {
	resp := &pb.SysAuthLoginCaptchaReply{}
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
