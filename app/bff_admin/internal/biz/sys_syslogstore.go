package biz

import (
	"context"
	sysV1 "fkratos/api/rpc_sys/v1"
	"fkratos/internal/errorx"

	pb "fkratos/api/bff_admin/v1"

	"github.com/jinzhu/copier"
)

// SysLogStore 日志-保存
func (s *SysUseCase) SysLogStore(ctx context.Context, req *pb.SysLogStoreReq) (*pb.SysLogStoreResp, error) {
	resp := new(pb.SysLogStoreResp)
	result, err := s.sysLogClient.SysLogStore(ctx, &sysV1.SysLogStoreReq{
		AdminID:   req.GetAdminID(),
		Username:  req.GetUsername(),
		Ip:        req.GetIp(),
		Uri:       req.GetUri(),
		UriDesc:   req.GetUriDesc(),
		Useragent: req.GetUseragent(),
		Req:       req.GetReq(),
		Resp:      req.GetResp(),
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
