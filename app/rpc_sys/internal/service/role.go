package service

import (
	"context"
	"fkratos/app/rpc_sys/internal/biz"

	pb "fkratos/api/rpc_sys/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type RoleService struct {
	pb.UnimplementedRoleServer

	log         *log.Helper
	roleUseCase *biz.RoleUseCase
}

func NewRoleService(logger log.Logger, roleUseCase *biz.RoleUseCase) *RoleService {
	l := log.NewHelper(log.With(logger, "module", "rpc_sys/service/role"))
	return &RoleService{
		log:         l,
		roleUseCase: roleUseCase,
	}
}

func (s *RoleService) SysRoleList(ctx context.Context, req *pb.SysRoleListReq) (*pb.SysRoleListResp, error) {
	return s.roleUseCase.SysRoleList(ctx, req)
}

func (s *RoleService) SysRoleInfo(ctx context.Context, req *pb.SysRoleInfoReq) (*pb.SysRoleInfoResp, error) {
	return s.roleUseCase.SysRoleInfo(ctx, req)
}

func (s *RoleService) SysRoleStore(ctx context.Context, req *pb.SysRoleStoreReq) (*pb.SysRoleStoreResp, error) {
	return s.roleUseCase.SysRoleStore(ctx, req)
}

func (s *RoleService) SysRoleDel(ctx context.Context, req *pb.SysRoleDelReq) (*pb.SysRoleDelResp, error) {
	return s.roleUseCase.SysRoleDel(ctx, req)
}
