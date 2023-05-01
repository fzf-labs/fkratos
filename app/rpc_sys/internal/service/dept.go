package service

import (
	"context"
	"fkratos/app/rpc_sys/internal/biz"

	pb "fkratos/api/rpc_sys/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type DeptService struct {
	pb.UnimplementedDeptServer

	log         *log.Helper
	deptUseCase *biz.DeptUseCase
}

func NewDeptService(logger log.Logger, deptUseCase *biz.DeptUseCase) *DeptService {
	l := log.NewHelper(log.With(logger, "module", "rpc_sys/service/dept"))
	return &DeptService{
		log:         l,
		deptUseCase: deptUseCase,
	}
}

func (s *DeptService) SysDeptList(ctx context.Context, req *pb.SysDeptListReq) (*pb.SysDeptListReply, error) {
	return s.deptUseCase.SysDeptList(ctx, req)
}
func (s *DeptService) SysDeptInfo(ctx context.Context, req *pb.SysDeptInfoReq) (*pb.SysDeptInfoReply, error) {
	return s.deptUseCase.SysDeptInfo(ctx, req)
}
func (s *DeptService) SysDeptStore(ctx context.Context, req *pb.SysDeptStoreReq) (*pb.SysDeptStoreReply, error) {
	return s.deptUseCase.SysDeptStore(ctx, req)
}
func (s *DeptService) SysDeptDel(ctx context.Context, req *pb.SysDeptDelReq) (*pb.SysDeptDelReply, error) {
	return s.deptUseCase.SysDeptDel(ctx, req)
}
