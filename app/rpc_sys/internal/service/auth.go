package service

import (
	"context"
	pb "fkratos/api/rpc_sys/v1"
	"fkratos/app/rpc_sys/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type AuthService struct {
	pb.UnimplementedAuthServer
	log *log.Helper

	authUseCase *biz.AuthUseCase
}

func NewAuthService(logger log.Logger, authUseCase *biz.AuthUseCase) *AuthService {
	l := log.NewHelper(log.With(logger, "module", "rpc_sys/service/auth"))
	return &AuthService{
		log:         l,
		authUseCase: authUseCase,
	}
}

func (s *AuthService) SysAuthLoginCaptcha(ctx context.Context, req *pb.SysAuthLoginCaptchaReq) (*pb.SysAuthLoginCaptchaReply, error) {
	return s.authUseCase.SysAuthLoginCaptcha(ctx, req)
}
func (s *AuthService) SysAuthLogin(ctx context.Context, req *pb.SysAuthLoginReq) (*pb.SysAuthLoginReply, error) {
	return s.authUseCase.SysAuthLogin(ctx, req)
}
func (s *AuthService) SysAuthLogout(ctx context.Context, req *pb.SysAuthLogoutReq) (*pb.SysAuthLogoutReply, error) {
	return s.authUseCase.SysAuthLogout(ctx, req)
}
