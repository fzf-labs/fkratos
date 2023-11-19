package biz

import (
	"context"
	sysV1 "fkratos/api/rpc_sys/v1"
	"fkratos/internal/errorx"

	pb "fkratos/api/bff_admin/v1"

	"github.com/jinzhu/copier"
)

// SysAPIDel API-删除
func (s *SysUseCase) SysAPIDel(ctx context.Context, req *pb.SysAPIDelReq) (*pb.SysAPIDelReply, error) {
	resp := new(pb.SysAPIDelReply)
	result, err := s.sysAPIClient.SysAPIDel(ctx, &sysV1.SysAPIDelReq{
		Ids: req.GetIds(),
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
