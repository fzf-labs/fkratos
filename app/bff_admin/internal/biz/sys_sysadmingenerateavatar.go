package biz

import (
	"context"
	sysV1 "fkratos/api/rpc_sys/v1"
	"fkratos/internal/errorx"
	"fkratos/internal/meta"

	pb "fkratos/api/bff_admin/v1"

	"github.com/jinzhu/copier"
)

// SysAdminGenerateAvatar 管理员-生成头像
func (s *SysUseCase) SysAdminGenerateAvatar(ctx context.Context, _ *pb.SysAdminGenerateAvatarReq) (*pb.SysAdminGenerateAvatarReply, error) {
	resp := new(pb.SysAdminGenerateAvatarReply)
	result, err := s.sysAdminClient.SysAdminGenerateAvatar(ctx, &sysV1.SysAdminGenerateAvatarReq{
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
