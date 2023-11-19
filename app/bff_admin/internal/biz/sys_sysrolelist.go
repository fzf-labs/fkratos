package biz

import (
	"context"
	sysV1 "fkratos/api/rpc_sys/v1"
	"fkratos/internal/errorx"

	pb "fkratos/api/bff_admin/v1"

	"github.com/jinzhu/copier"
)

// SysRoleList 角色-列表
func (s *SysUseCase) SysRoleList(ctx context.Context, _ *pb.SysRoleListReq) (*pb.SysRoleListResp, error) {
	resp := new(pb.SysRoleListResp)
	result, err := s.sysRoleClient.SysRoleList(ctx, &sysV1.SysRoleListReq{})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(resp, result)
	if err != nil {
		return nil, errorx.DataFormattingError.WithError(err).Err()
	}
	return resp, nil
}
