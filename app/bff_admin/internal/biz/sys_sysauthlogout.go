package biz

import (
	"context"
	sysV1 "fkratos/api/rpc_sys/v1"
	"fkratos/internal/errorx"
	"fkratos/internal/meta"

	pb "fkratos/api/bff_admin/v1"

	"github.com/jinzhu/copier"
)

// SysAuthLogout Auth-退出
func (s *SysUseCase) SysAuthLogout(ctx context.Context, _ *pb.SysAuthLogoutReq) (*pb.SysAuthLogoutReply, error) {
	resp := new(pb.SysAuthLogoutReply)
	result, err := s.sysAuthClient.SysAuthLogout(ctx, &sysV1.SysAuthLogoutReq{
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
