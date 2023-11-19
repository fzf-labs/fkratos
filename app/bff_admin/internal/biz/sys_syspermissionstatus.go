package biz

import (
	"context"
	sysV1 "fkratos/api/rpc_sys/v1"
	"fkratos/internal/errorx"

	pb "fkratos/api/bff_admin/v1"

	"github.com/jinzhu/copier"
)

// SysPermissionStatus 权限-修改状态
func (s *SysUseCase) SysPermissionStatus(ctx context.Context, req *pb.SysPermissionStatusReq) (*pb.SysPermissionStatusResp, error) {
	resp := new(pb.SysPermissionStatusResp)
	result, err := s.sysPermissionClient.SysPermissionStatus(ctx, &sysV1.SysPermissionStatusReq{
		Id:     req.GetId(),
		Status: req.GetStatus(),
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(resp, result)
	if err != nil {
		return nil, errorx.DataFormattingError.WithError(err).Err()
	}
	return resp, nil
}
