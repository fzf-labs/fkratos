package biz

import (
	"context"
	"fkratos/internal/errorx"

	pb "fkratos/api/rpc_sys/v1"

	"github.com/fzf-labs/fpkg/util/jsonutil"
)

// SysAdminInfo 管理员-个人信息
func (s *SysAdminUseCase) SysAdminInfo(ctx context.Context, req *pb.SysAdminInfoReq) (*pb.SysAdminInfoReply, error) {
	resp := &pb.SysAdminInfoReply{}
	sysAdmin, err := s.sysAdminRepo.FindOneCacheByID(ctx, req.GetAdminId())
	if err != nil {
		return nil, err
	}
	if sysAdmin == nil {
		return resp, nil
	}
	roleIds := make([]string, 0)
	if sysAdmin.RoleIds.String() != "" {
		err = jsonutil.Unmarshal(sysAdmin.RoleIds, &roleIds)
		if err != nil {
			return nil, errorx.DataFormattingError.WithError(err).Err()
		}
	}
	resp.Info = &pb.SysAdminInfo{
		Id:       sysAdmin.ID,
		Username: sysAdmin.Username,
		Nickname: sysAdmin.Nickname,
		Avatar:   sysAdmin.Avatar,
		Gender:   int32(sysAdmin.Gender),
		Email:    sysAdmin.Email,
		Mobile:   sysAdmin.Mobile,
		JobID:    sysAdmin.JobID,
		DeptID:   sysAdmin.DeptID,
		RoleIds:  roleIds,
		Motto:    sysAdmin.Motto,
	}
	return resp, nil
}
