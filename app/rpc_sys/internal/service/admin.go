package service

import (
	"context"
	"fkratos/api/paginator"
	pb "fkratos/api/rpc_sys/v1"
	"fkratos/app/rpc_sys/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type AdminService struct {
	pb.UnimplementedAdminServer
	log          *log.Helper
	adminUseCase *biz.AdminUseCase
}

func NewAdminService(logger log.Logger, adminUseCase *biz.AdminUseCase) *AdminService {
	l := log.NewHelper(log.With(logger, "module", "rpc_sys/service/admin"))
	return &AdminService{
		log:          l,
		adminUseCase: adminUseCase,
	}
}

func (s *AdminService) SysAdminInfo(ctx context.Context, req *pb.SysAdminInfoReq) (*pb.SysAdminInfoReply, error) {
	return s.adminUseCase.SysAdminInfo(ctx, req)
}
func (s *AdminService) SysAdminInfoUpdate(ctx context.Context, req *pb.SysAdminInfoUpdateReq) (*pb.SysAdminInfoUpdateReply, error) {
	return s.adminUseCase.SysAdminInfoUpdate(ctx, req)
}
func (s *AdminService) SysAdminGenerateAvatar(ctx context.Context, req *pb.SysAdminGenerateAvatarReq) (*pb.SysAdminGenerateAvatarReply, error) {
	return s.adminUseCase.SysAdminGenerateAvatar(ctx, req)
}
func (s *AdminService) SysManageList(ctx context.Context, req *paginator.PaginatorReq) (*pb.SysManageListReply, error) {
	return s.adminUseCase.SysManageList(ctx, req)
}
func (s *AdminService) SysManageInfo(ctx context.Context, req *pb.SysManageInfoReq) (*pb.SysManageInfoReply, error) {
	return s.adminUseCase.SysManageInfo(ctx, req)
}
func (s *AdminService) SysManageStore(ctx context.Context, req *pb.SysManageStoreReq) (*pb.SysManageStoreReply, error) {
	return s.adminUseCase.SysManageStore(ctx, req)
}
func (s *AdminService) SysManageDel(ctx context.Context, req *pb.SysManageDelReq) (*pb.SysManageDelReply, error) {
	return s.adminUseCase.SysManageDel(ctx, req)
}
func (s *AdminService) SysAdminPermission(ctx context.Context, req *pb.SysAdminPermissionReq) (*pb.SysAdminPermissionReply, error) {
	return s.adminUseCase.SysAdminPermission(ctx, req)
}
