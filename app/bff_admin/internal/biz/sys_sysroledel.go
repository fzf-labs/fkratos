package biz

import (
	"context"
	sysV1 "fkratos/api/rpc_sys/v1"
	"fkratos/internal/errorx"

	pb "fkratos/api/bff_admin/v1"

	"github.com/jinzhu/copier"
)

// SysRoleDel 角色-删除

func (s *SysUseCase) SysRoleDel(ctx context.Context, req *pb.SysRoleDelReq) (*pb.SysRoleDelResp, error) {
	resp := new(pb.SysRoleDelResp)
	result, err := s.sysRoleClient.SysRoleDel(ctx, &sysV1.SysRoleDelReq{
		Ids: req.GetIds(),
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(resp, result)
	if err != nil {
		return nil, errorx.DataFormattingError.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return resp, nil
}
