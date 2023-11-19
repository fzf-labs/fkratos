package biz

import (
	"context"
	sysV1 "fkratos/api/rpc_sys/v1"
	"fkratos/internal/errorx"

	pb "fkratos/api/bff_admin/v1"

	"github.com/jinzhu/copier"
)

// SysManageList 管理员-列表
func (s *SysUseCase) SysManageList(ctx context.Context, req *pb.SysManageListReq) (*pb.SysManageListReply, error) {
	resp := new(pb.SysManageListReply)
	result, err := s.sysAdminClient.SysManageList(ctx, &sysV1.SysManageListReq{
		Paginator: req.GetPaginator(),
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
