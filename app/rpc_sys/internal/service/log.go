package service

import (
	"context"
	"fkratos/api/common"
	"fkratos/app/rpc_sys/internal/biz"

	pb "fkratos/api/rpc_sys/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type LogService struct {
	pb.UnimplementedLogServer

	log        *log.Helper
	logUseCase *biz.LogUseCase
}

func NewLogService(logger log.Logger, logUseCase *biz.LogUseCase) *LogService {
	l := log.NewHelper(log.With(logger, "module", "rpc_sys/service/job"))
	return &LogService{
		log:        l,
		logUseCase: logUseCase,
	}
}

func (s *LogService) SysLogList(ctx context.Context, req *common.SearchListReq) (*pb.SysLogListResp, error) {
	return s.logUseCase.SysLogList(ctx, req)
}

func (s *LogService) SysLogInfo(ctx context.Context, req *pb.SysLogInfoReq) (*pb.SysLogInfoResp, error) {
	return s.logUseCase.SysLogInfo(ctx, req)
}
