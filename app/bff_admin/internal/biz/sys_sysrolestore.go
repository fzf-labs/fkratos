package biz

import (
	"context"
	sysV1 "fkratos/api/rpc_sys/v1"
	"fkratos/internal/errorx"

	pb "fkratos/api/bff_admin/v1"

	"github.com/jinzhu/copier"
)

// SysRoleStore 角色-保存
func (s *SysUseCase) SysRoleStore(ctx context.Context, req *pb.SysRoleStoreReq) (*pb.SysRoleStoreResp, error) {
	resp := new(pb.SysRoleStoreResp)
	result, err := s.sysRoleClient.SysRoleStore(ctx, &sysV1.SysRoleStoreReq{
		Id:            req.GetId(),
		Pid:           req.GetPid(),
		Name:          req.GetName(),
		Remark:        req.GetRemark(),
		Status:        req.GetStatus(),
		PermissionIds: req.GetPermissionIds(),
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
