package data

import (
	"context"
	v1 "fkratos/api/demo/v1"
	"fkratos/app/demo/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

func NewDemoRepo(data *Data, logger log.Logger) biz.DemoRepo {
	l := log.NewHelper(log.With(logger, "module", "Demo/repo/Demo-service"))
	return &DemoRepo{
		data: data,
		log:  l,
	}
}

type DemoRepo struct {
	data *Data
	log  *log.Helper
}

func (d *DemoRepo) CreateDemo(ctx context.Context, req *v1.CreateDemoRequest) (*v1.CreateDemoReply, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DemoRepo) UpdateDemo(ctx context.Context, req *v1.UpdateDemoRequest) (*v1.UpdateDemoReply, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DemoRepo) DeleteDemo(ctx context.Context, req *v1.DeleteDemoRequest) (*v1.DeleteDemoReply, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DemoRepo) GetDemo(ctx context.Context, req *v1.GetDemoRequest) (*v1.GetDemoReply, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DemoRepo) ListDemo(ctx context.Context, req *v1.ListDemoRequest) (*v1.ListDemoReply, error) {
	//TODO implement me
	panic("implement me")
}
