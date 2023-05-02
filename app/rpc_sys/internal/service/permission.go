package service

import (
	"context"
	"fkratos/app/rpc_sys/internal/biz"

	pb "fkratos/api/rpc_sys/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type PermissionService struct {
	pb.UnimplementedPermissionServer

	log               *log.Helper
	permissionUseCase *biz.PermissionUseCase
}

func NewPermissionService(logger log.Logger, permissionUseCase *biz.PermissionUseCase) *PermissionService {
	l := log.NewHelper(log.With(logger, "module", "rpc_sys/service/role"))
	return &PermissionService{
		log:               l,
		permissionUseCase: permissionUseCase,
	}
}

func (s *PermissionService) SysPermissionList(ctx context.Context, req *pb.SysPermissionListReq) (*pb.SysPermissionListResp, error) {
	return s.permissionUseCase.SysPermissionList(ctx, req)
}
func (s *PermissionService) SysPermissionInfo(ctx context.Context, req *pb.SysPermissionInfoReq) (*pb.SysPermissionInfoResp, error) {
	return s.permissionUseCase.SysPermissionInfo(ctx, req)
}
func (s *PermissionService) SysPermissionStore(ctx context.Context, req *pb.SysPermissionStoreReq) (*pb.SysPermissionStoreResp, error) {
	return s.permissionUseCase.SysPermissionStore(ctx, req)
}
func (s *PermissionService) SysPermissionDel(ctx context.Context, req *pb.SysPermissionDelReq) (*pb.SysPermissionDelResp, error) {
	return s.permissionUseCase.SysPermissionDel(ctx, req)
}
func (s *PermissionService) SysPermissionStatus(ctx context.Context, req *pb.SysPermissionStatusReq) (*pb.SysPermissionStatusResp, error) {
	return s.permissionUseCase.SysPermissionStatus(ctx, req)
}
