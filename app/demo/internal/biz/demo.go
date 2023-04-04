package biz

import (
	"context"
	v1 "fkratos/api/demo/v1"
	"github.com/go-kratos/kratos/v2/log"
)

type DemoRepo interface {
	CreateDemo(ctx context.Context, req *v1.CreateDemoRequest) (*v1.CreateDemoReply, error)
	UpdateDemo(ctx context.Context, req *v1.UpdateDemoRequest) (*v1.UpdateDemoReply, error)
	DeleteDemo(ctx context.Context, req *v1.DeleteDemoRequest) (*v1.DeleteDemoReply, error)
	GetDemo(ctx context.Context, req *v1.GetDemoRequest) (*v1.GetDemoReply, error)
	ListDemo(ctx context.Context, req *v1.ListDemoRequest) (*v1.ListDemoReply, error)
}
type DemoUseCase struct {
	repo DemoRepo
	log  *log.Helper
}

func NewDemoUseCase(repo DemoRepo, logger log.Logger) *DemoUseCase {
	l := log.NewHelper(log.With(logger, "module", "demo/usecase/demo-service"))
	return &DemoUseCase{
		repo: repo,
		log:  l,
	}
}

func (d *DemoUseCase) CreateDemo(ctx context.Context, req *v1.CreateDemoRequest) (*v1.CreateDemoReply, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DemoUseCase) UpdateDemo(ctx context.Context, req *v1.UpdateDemoRequest) (*v1.UpdateDemoReply, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DemoUseCase) DeleteDemo(ctx context.Context, req *v1.DeleteDemoRequest) (*v1.DeleteDemoReply, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DemoUseCase) GetDemo(ctx context.Context, req *v1.GetDemoRequest) (*v1.GetDemoReply, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DemoUseCase) ListDemo(ctx context.Context, req *v1.ListDemoRequest) (*v1.ListDemoReply, error) {
	//TODO implement me
	panic("implement me")
}
