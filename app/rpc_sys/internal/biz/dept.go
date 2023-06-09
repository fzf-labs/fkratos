package biz

import (
	"context"
	v1 "fkratos/api/rpc_sys/v1"

	"github.com/go-kratos/kratos/v2/log"
)

func NewDeptUseCase(logger log.Logger, sysDeptRepo SysDeptRepo) *DeptUseCase {
	l := log.NewHelper(log.With(logger, "module", "rpc_user/biz/dept"))
	return &DeptUseCase{
		log:         l,
		sysDeptRepo: sysDeptRepo,
	}
}

type DeptUseCase struct {
	log         *log.Helper
	sysDeptRepo SysDeptRepo
}

func (d *DeptUseCase) SysDeptList(ctx context.Context, req *v1.SysDeptListReq) (*v1.SysDeptListReply, error) {
	resp := new(v1.SysDeptListReply)
	list, err := d.sysDeptRepo.SysDeptList(ctx)
	if err != nil {
		return nil, err
	}
	resp.List = list
	return resp, nil
}

func (d *DeptUseCase) SysDeptInfo(ctx context.Context, req *v1.SysDeptInfoReq) (*v1.SysDeptInfoReply, error) {
	resp := new(v1.SysDeptInfoReply)
	info, err := d.sysDeptRepo.SysDeptInfoById(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	resp.Info = info
	return resp, nil
}

func (d *DeptUseCase) SysDeptStore(ctx context.Context, req *v1.SysDeptStoreReq) (*v1.SysDeptStoreReply, error) {
	resp := new(v1.SysDeptStoreReply)
	_, err := d.sysDeptRepo.SysDeptStore(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DeptUseCase) SysDeptDel(ctx context.Context, req *v1.SysDeptDelReq) (*v1.SysDeptDelReply, error) {
	resp := new(v1.SysDeptDelReply)
	err := d.sysDeptRepo.SysDeptDelByIds(ctx, req.GetIds())
	if err != nil {
		return nil, err
	}
	return resp, nil
}
