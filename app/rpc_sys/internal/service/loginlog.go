package service

import (
	"context"
	"fkratos/api/common"

	pb "fkratos/api/rpc_sys/v1"
)

type LoginLogService struct {
	pb.UnimplementedLoginLogServer
}

func NewLoginLogService() *LoginLogService {
	return &LoginLogService{}
}

func (s *LoginLogService) SysLoginLogList(ctx context.Context, req *common.SearchListReq) (*pb.SysLoginLogListResp, error) {
	return &pb.SysLoginLogListResp{}, nil
}
func (s *LoginLogService) SysLoginLogInfo(ctx context.Context, req *pb.SysLoginLogInfoReq) (*pb.SysLoginLogInfoResp, error) {
	return &pb.SysLoginLogInfoResp{}, nil
}
func (s *LoginLogService) SysLoginLogStore(ctx context.Context, req *pb.SysLoginLogStoreReq) (*pb.SysLoginLogStoreResp, error) {
	return &pb.SysLoginLogStoreResp{}, nil
}
