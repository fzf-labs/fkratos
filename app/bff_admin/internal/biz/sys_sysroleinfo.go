package biz

import (
	"context"
	sysV1 "fkratos/api/rpc_sys/v1"
	"fkratos/internal/errorx"

	pb "fkratos/api/bff_admin/v1"

	"github.com/jinzhu/copier"
)

// SysRoleInfo 角色-信息
func (s *SysUseCase) SysRoleInfo(ctx context.Context, req *pb.SysRoleInfoReq) (*pb.SysRoleInfoResp, error) {
	resp := new(pb.SysRoleInfoResp)
	result, err := s.sysRoleClient.SysRoleInfo(ctx, &sysV1.SysRoleInfoReq{
		Id: req.GetId(),
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
