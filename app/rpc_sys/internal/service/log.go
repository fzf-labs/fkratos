package service

import (
	"context"
	"encoding/json"
	"fkratos/api/common"
	pb "fkratos/api/rpc_sys/v1"
	"fkratos/app/rpc_sys/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type LogService struct {
	pb.UnimplementedLogServer

	log        *log.Helper
	logUseCase *biz.LogUseCase
}

func NewLogService(logger log.Logger, logUseCase *biz.LogUseCase) *LogService {
	l := log.NewHelper(log.With(logger, "module", "rpc_sys/service/log"))
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

func (s *LogService) SysLogStore(ctx context.Context, req *pb.SysLogStoreReq) (*pb.SysLogStoreResp, error) {
	err := s.SysLogStoreProducer(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.SysLogStoreResp{}, nil
}

// SysLogStoreProducer 日志生产者
func (s *LogService) SysLogStoreProducer(ctx context.Context, req *pb.SysLogStoreReq) error {
	err := s.logUseCase.SysLogStoreProducer(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

// SysLogStoreConsumer 日志消费者
func (s *LogService) SysLogStoreConsumer(payload []byte) error {
	var params *pb.SysLogStoreReq
	err := json.Unmarshal(payload, params)
	if err != nil {
		return err
	}
	_, err = s.logUseCase.SysLogStore(context.Background(), params)
	if err != nil {
		return err
	}
	return nil
}
