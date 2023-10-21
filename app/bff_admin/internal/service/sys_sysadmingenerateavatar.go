package service

import (
	"context"

	pb "fkratos/api/bff_admin/v1"
)

func (s *SysService) SysAdminGenerateAvatar(ctx context.Context, req *pb.SysAdminGenerateAvatarReq) (*pb.SysAdminGenerateAvatarReply, error) {
	return s.sysUseCase.SysAdminGenerateAvatar(ctx, req)
}
