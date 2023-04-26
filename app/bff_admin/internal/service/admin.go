package service

import (
	"context"
	userV1 "fkratos/api/rpc_user/v1"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"

	pb "fkratos/api/bff_admin/v1"
)

type AdminService struct {
	pb.UnimplementedAdminServer
	log        *log.Helper
	userClient userV1.UserClient
}

func NewAdminService(logger log.Logger, userClient userV1.UserClient) *AdminService {
	l := log.NewHelper(log.With(logger, "module", "user/bff_admin/admin-service"))
	return &AdminService{
		log:        l,
		userClient: userClient,
	}
}

func (s *AdminService) CreateAdmin(ctx context.Context, req *pb.CreateAdminReq) (*pb.CreateAdminReply, error) {
	user, err := s.userClient.CreateUser(ctx, &userV1.CreateUserReq{})
	if err != nil {
		return nil, err
	}
	fmt.Println(user)
	return &pb.CreateAdminReply{}, nil
}
func (s *AdminService) UpdateAdmin(ctx context.Context, req *pb.UpdateAdminReq) (*pb.UpdateAdminReply, error) {
	return &pb.UpdateAdminReply{}, nil
}
func (s *AdminService) DeleteAdmin(ctx context.Context, req *pb.DeleteAdminReq) (*pb.DeleteAdminReply, error) {
	return &pb.DeleteAdminReply{}, nil
}
func (s *AdminService) GetAdmin(ctx context.Context, req *pb.GetAdminReq) (*pb.GetAdminReply, error) {
	return &pb.GetAdminReply{}, nil
}
func (s *AdminService) ListAdmin(ctx context.Context, req *pb.ListAdminReq) (*pb.ListAdminReply, error) {
	return &pb.ListAdminReply{}, nil
}
