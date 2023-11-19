package biz

import (
	"context"
	sysV1 "fkratos/api/rpc_sys/v1"
	"fkratos/internal/errorx"

	pb "fkratos/api/bff_admin/v1"

	"github.com/jinzhu/copier"
)

// SysAPIStore API-保存
func (s *SysUseCase) SysAPIStore(ctx context.Context, req *pb.SysAPIStoreReq) (*pb.SysAPIStoreReply, error) {
	resp := new(pb.SysAPIStoreReply)
	result, err := s.sysAPIClient.SysAPIStore(ctx, &sysV1.SysAPIStoreReq{
		Id:           req.GetId(),
		PermissionID: req.GetPermissionID(),
		Method:       req.GetMethod(),
		Path:         req.GetPath(),
		Desc:         req.GetDesc(),
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
