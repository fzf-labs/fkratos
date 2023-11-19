package biz

import (
	"context"
	sysV1 "fkratos/api/rpc_sys/v1"
	"fkratos/internal/errorx"

	pb "fkratos/api/bff_admin/v1"

	"github.com/jinzhu/copier"
)

// SysManageInfo 管理员-信息
func (s *SysUseCase) SysManageInfo(ctx context.Context, req *pb.SysManageInfoReq) (*pb.SysManageInfoReply, error) {
	resp := new(pb.SysManageInfoReply)
	result, err := s.sysAdminClient.SysManageInfo(ctx, &sysV1.SysManageInfoReq{
		AdminId: req.GetAdminId(),
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
