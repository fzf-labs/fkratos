package biz

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"

	"github.com/fzf-labs/fpkg/third_api/avatar"
)

// SysAdminGenerateAvatar 管理员-生成头像
func (s *SysAdminUseCase) SysAdminGenerateAvatar(_ context.Context, _ *pb.SysAdminGenerateAvatarReq) (*pb.SysAdminGenerateAvatarReply, error) {
	return &pb.SysAdminGenerateAvatarReply{AvatarUrl: avatar.URL()}, nil
}
