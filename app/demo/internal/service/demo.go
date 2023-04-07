package service

import (
	"context"
	"fkratos/app/demo/internal/biz"
	"github.com/go-kratos/kratos/v2/log"

	pb "fkratos/api/demo/v1"
)

type DemoService struct {
	pb.UnimplementedDemoServer

	log *log.Helper

	demoUseCase *biz.DemoUseCase
}

func NewDemoService(logger log.Logger, demoUseCase *biz.DemoUseCase) *DemoService {
	l := log.NewHelper(log.With(logger, "module", "user/service/demo-service"))
	return &DemoService{
		log:         l,
		demoUseCase: demoUseCase,
	}
}

func (s *DemoService) CreateDemo(ctx context.Context, req *pb.CreateDemoReq) (*pb.CreateDemoReply, error) {
	return &pb.CreateDemoReply{}, nil
}
func (s *DemoService) UpdateDemo(ctx context.Context, req *pb.UpdateDemoReq) (*pb.UpdateDemoReply, error) {
	return &pb.UpdateDemoReply{}, nil
}
func (s *DemoService) DeleteDemo(ctx context.Context, req *pb.DeleteDemoReq) (*pb.DeleteDemoReply, error) {
	return &pb.DeleteDemoReply{}, nil
}
func (s *DemoService) GetDemo(ctx context.Context, req *pb.GetDemoReq) (*pb.GetDemoReply, error) {
	return &pb.GetDemoReply{}, nil
}
func (s *DemoService) ListDemo(ctx context.Context, req *pb.ListDemoReq) (*pb.ListDemoReply, error) {
	return &pb.ListDemoReply{}, nil
}
