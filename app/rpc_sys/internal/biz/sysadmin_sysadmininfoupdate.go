package biz

import (
	"context"
	"fkratos/internal/errorx"

	pb "fkratos/api/rpc_sys/v1"

	"github.com/fzf-labs/fpkg/crypt"
)

// SysAdminInfoUpdate 管理员-个人信息更新
func (s *SysAdminUseCase) SysAdminInfoUpdate(ctx context.Context, req *pb.SysAdminInfoUpdateReq) (*pb.SysAdminInfoUpdateReply, error) {
	resp := new(pb.SysAdminInfoUpdateReply)
	sysAdminInfo, err := s.sysAdminRepo.FindOneCacheByID(ctx, req.GetAdminId())
	if err != nil {
		return nil, errorx.DataSQLErr.WithError(err).Err()
	}
	if sysAdminInfo == nil {
		return nil, errorx.AccountNotExist.Err()
	}
	var pwd string
	if req.Password != "" {
		pwd, err = crypt.Encrypt(req.Password + sysAdminInfo.Salt)
		if err != nil {
			return resp, errorx.DataProcessingError.WithError(err).Err()
		}
	}
	sysAdminInfo.Password = pwd
	sysAdminInfo.Nickname = req.Nickname
	sysAdminInfo.Email = req.Email
	sysAdminInfo.Mobile = req.Mobile
	sysAdminInfo.Motto = req.Motto
	err = s.sysAdminRepo.UpdateOne(ctx, sysAdminInfo)
	if err != nil {
		return nil, errorx.DataSQLErr.WithError(err).Err()
	}
	return resp, nil
}
