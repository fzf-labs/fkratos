package service

import (
	"context"
	"fkratos/api/rpc/admin/v1"
	"fkratos/app/bff/admin/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type DemoService struct {
	v1.UnimplementedDemoServer

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

func (s *DemoService) CreateDemo(ctx context.Context, req *v1.CreateDemoReq) (*v1.CreateDemoReply, error) {
	return &v1.CreateDemoReply{}, nil
}
func (s *DemoService) UpdateDemo(ctx context.Context, req *v1.UpdateDemoReq) (*v1.UpdateDemoReply, error) {
	return &v1.UpdateDemoReply{}, nil
}
func (s *DemoService) DeleteDemo(ctx context.Context, req *v1.DeleteDemoReq) (*v1.DeleteDemoReply, error) {
	return &v1.DeleteDemoReply{}, nil
}
func (s *DemoService) GetDemo(ctx context.Context, req *v1.GetDemoReq) (*v1.GetDemoReply, error) {
	return &v1.GetDemoReply{}, nil
}
func (s *DemoService) ListDemo(ctx context.Context, req *v1.ListDemoReq) (*v1.ListDemoReply, error) {
	return &v1.ListDemoReply{}, nil
}
