package biz

import (
	"context"
	sysV1 "fkratos/api/rpc_sys/v1"
	"fkratos/internal/errorx"
	"fkratos/internal/meta"

	pb "fkratos/api/bff_admin/v1"

	"github.com/jinzhu/copier"
)

// SysAdminInfo 管理员-个人信息
func (s *SysUseCase) SysAdminInfo(ctx context.Context, _ *pb.SysAdminInfoReq) (*pb.SysAdminInfoReply, error) {
	resp := new(pb.SysAdminInfoReply)
	result, err := s.sysAdminClient.SysAdminInfo(ctx, &sysV1.SysAdminInfoReq{
		AdminId: meta.GetAdminIDFromClient(ctx),
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
