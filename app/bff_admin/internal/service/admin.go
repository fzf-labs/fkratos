package service

import (
	"context"

	pb "fkratos/api/bff_admin/v1"
)

type AdminService struct {
	pb.UnimplementedAdminServer
}

func NewAdminService() *AdminService {
	return &AdminService{}
}

func (s *AdminService) CreateAdmin(ctx context.Context, req *pb.CreateAdminReq) (*pb.CreateAdminReply, error) {
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
