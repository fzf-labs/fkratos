package service

import (
	"context"
	pb "fkratos/api/rpc_sys/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type AuthService struct {
	pb.UnimplementedAuthServer

	log *log.Helper
}

func NewAuthService(logger log.Logger) *AuthService {
	l := log.NewHelper(log.With(logger, "module", "user/service/user-service"))
	return &AuthService{
		log: l,
	}
}

func (s *AuthService) SysAuthLoginCaptcha(ctx context.Context, req *pb.SysAuthLoginCaptchaReq) (*pb.SysAuthLoginCaptchaReply, error) {
	return &pb.SysAuthLoginCaptchaReply{}, nil
}
func (s *AuthService) SysAuthLogin(ctx context.Context, req *pb.SysAuthLoginReq) (*pb.SysAuthLoginReply, error) {
	return &pb.SysAuthLoginReply{}, nil
}
func (s *AuthService) SysAuthLogout(ctx context.Context, req *pb.SysAuthLogoutReq) (*pb.SysAuthLogoutReply, error) {
	return &pb.SysAuthLogoutReply{}, nil
}
