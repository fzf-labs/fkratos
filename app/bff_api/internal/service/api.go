package service

import (
	"context"

	pb "fkratos/api/bff_api/v1"
)

type ApiService struct {
	pb.UnimplementedApiServer
}

func NewApiService() *ApiService {
	return &ApiService{}
}

func (s *ApiService) CreateApi(ctx context.Context, req *pb.CreateApiReq) (*pb.CreateApiReply, error) {
	return &pb.CreateApiReply{}, nil
}
func (s *ApiService) UpdateApi(ctx context.Context, req *pb.UpdateApiReq) (*pb.UpdateApiReply, error) {
	return &pb.UpdateApiReply{}, nil
}
func (s *ApiService) DeleteApi(ctx context.Context, req *pb.DeleteApiReq) (*pb.DeleteApiReply, error) {
	return &pb.DeleteApiReply{}, nil
}
func (s *ApiService) GetApi(ctx context.Context, req *pb.GetApiReq) (*pb.GetApiReply, error) {
	return &pb.GetApiReply{}, nil
}
func (s *ApiService) ListApi(ctx context.Context, req *pb.ListApiReq) (*pb.ListApiReply, error) {
	return &pb.ListApiReply{}, nil
}
