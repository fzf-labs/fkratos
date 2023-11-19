package biz

import (
	"context"
	sysV1 "fkratos/api/rpc_sys/v1"
	"fkratos/internal/errorx"

	pb "fkratos/api/bff_admin/v1"

	"github.com/jinzhu/copier"
)

// SysPermissionList 权限-列表
func (s *SysUseCase) SysPermissionList(ctx context.Context, _ *pb.SysPermissionListReq) (*pb.SysPermissionListResp, error) {
	resp := new(pb.SysPermissionListResp)
	result, err := s.sysPermissionClient.SysPermissionList(ctx, &sysV1.SysPermissionListReq{})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(resp, result)
	if err != nil {
		return nil, errorx.DataFormattingError.WithError(err).Err()
	}
	return resp, nil
}
