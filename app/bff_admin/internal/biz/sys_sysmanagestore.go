package biz

import (
	"context"
	sysV1 "fkratos/api/rpc_sys/v1"
	"fkratos/internal/errorx"

	pb "fkratos/api/bff_admin/v1"

	"github.com/jinzhu/copier"
)

// SysManageStore 管理员-保存
func (s *SysUseCase) SysManageStore(ctx context.Context, req *pb.SysManageStoreReq) (*pb.SysManageStoreReply, error) {
	resp := new(pb.SysManageStoreReply)
	result, err := s.sysAdminClient.SysManageStore(ctx, &sysV1.SysManageStoreReq{
		Id:       req.GetId(),
		Username: req.GetUsername(),
		Nickname: req.GetNickname(),
		Password: req.GetPassword(),
		Avatar:   req.GetAvatar(),
		Gender:   req.GetGender(),
		Email:    req.GetEmail(),
		Mobile:   req.GetMobile(),
		JobID:    req.GetJobID(),
		DeptID:   req.GetDeptID(),
		RoleIds:  req.GetRoleIds(),
		Motto:    req.GetMotto(),
		Status:   req.GetStatus(),
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
