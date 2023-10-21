package service

import (
	"context"

	pb "fkratos/api/bff_admin/v1"
)

func (s *SysService) SysPermissionInfo(ctx context.Context, req *pb.SysPermissionInfoReq) (*pb.SysPermissionInfoResp, error) {
	return s.sysUseCase.SysPermissionInfo(ctx, req)
}
