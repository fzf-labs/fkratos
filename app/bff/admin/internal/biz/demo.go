package biz

import (
	"context"
	"fkratos/api/rpc/admin/v1"
	"github.com/go-kratos/kratos/v2/log"
)

type DemoRepo interface {
	CreateDemo(ctx context.Context, req *v1.CreateDemoReq) (*v1.CreateDemoReply, error)
	UpdateDemo(ctx context.Context, req *v1.UpdateDemoReq) (*v1.UpdateDemoReply, error)
	DeleteDemo(ctx context.Context, req *v1.DeleteDemoReq) (*v1.DeleteDemoReply, error)
	GetDemo(ctx context.Context, req *v1.GetDemoReq) (*v1.GetDemoReply, error)
	ListDemo(ctx context.Context, req *v1.ListDemoReq) (*v1.ListDemoReply, error)
}
type DemoUseCase struct {
	repo DemoRepo
	log  *log.Helper
}

func NewDemoUseCase(repo DemoRepo, logger log.Logger) *DemoUseCase {
	l := log.NewHelper(log.With(logger, "module", "admin/usecase/demo-service"))
	return &DemoUseCase{
		repo: repo,
		log:  l,
	}
}

func (d *DemoUseCase) CreateDemo(ctx context.Context, req *v1.CreateDemoReq) (*v1.CreateDemoReply, error) {
	return nil, nil
}

func (d *DemoUseCase) UpdateDemo(ctx context.Context, req *v1.UpdateDemoReq) (*v1.UpdateDemoReply, error) {
	return nil, nil
}

func (d *DemoUseCase) DeleteDemo(ctx context.Context, req *v1.DeleteDemoReq) (*v1.DeleteDemoReply, error) {
	return nil, nil
}

func (d *DemoUseCase) GetDemo(ctx context.Context, req *v1.GetDemoReq) (*v1.GetDemoReply, error) {
	return nil, nil
}

func (d *DemoUseCase) ListDemo(ctx context.Context, req *v1.ListDemoReq) (*v1.ListDemoReply, error) {
	return nil, nil
}
