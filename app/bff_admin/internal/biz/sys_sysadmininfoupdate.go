package biz

import (
	"context"
	sysV1 "fkratos/api/rpc_sys/v1"
	"fkratos/internal/errorx"
	"fkratos/internal/meta"

	pb "fkratos/api/bff_admin/v1"

	"github.com/jinzhu/copier"
)

// SysAdminInfoUpdate 管理员-个人信息更新
func (s *SysUseCase) SysAdminInfoUpdate(ctx context.Context, req *pb.SysAdminInfoUpdateReq) (*pb.SysAdminInfoUpdateReply, error) {
	resp := new(pb.SysAdminInfoUpdateReply)
	result, err := s.sysAdminClient.SysAdminInfoUpdate(ctx, &sysV1.SysAdminInfoUpdateReq{
		AdminId:  meta.GetAdminIDFromClient(ctx),
		Nickname: req.GetNickname(),
		Email:    req.GetEmail(),
		Mobile:   req.GetMobile(),
		Motto:    req.GetMotto(),
		Password: req.GetPassword(),
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
