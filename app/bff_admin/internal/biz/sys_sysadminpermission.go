package biz

import (
	"context"
	sysV1 "fkratos/api/rpc_sys/v1"
	"fkratos/internal/errorx"
	"fkratos/internal/meta"

	pb "fkratos/api/bff_admin/v1"

	"github.com/jinzhu/copier"
)

// SysAdminPermission 管理员-查询权限
func (s *SysUseCase) SysAdminPermission(ctx context.Context, _ *pb.SysAdminPermissionReq) (*pb.SysAdminPermissionReply, error) {
	resp := new(pb.SysAdminPermissionReply)
	result, err := s.sysAdminClient.SysAdminPermission(ctx, &sysV1.SysAdminPermissionReq{
		AdminId: meta.GetAdminIDFromClient(ctx),
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
