package service

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

func (s *SysAdminService) SysAdminGenerateAvatar(ctx context.Context, req *pb.SysAdminGenerateAvatarReq) (*pb.SysAdminGenerateAvatarReply, error) {
	return s.sysAdminUseCase.SysAdminGenerateAvatar(ctx, req)
}
