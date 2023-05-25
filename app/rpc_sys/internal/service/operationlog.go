package service

import (
	"context"
	"encoding/json"
	"fkratos/api/common"
	"fkratos/app/rpc_sys/internal/biz"

	pb "fkratos/api/rpc_sys/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type OperationLogService struct {
	pb.UnimplementedOperationLogServer

	log                 *log.Helper
	operationLogUseCase *biz.OperationLogUseCase
}

func NewOperationLogService(logger log.Logger, operationLogUseCase *biz.OperationLogUseCase) *OperationLogService {
	l := log.NewHelper(log.With(logger, "module", "rpc_sys/service/job"))
	return &OperationLogService{
		log:                 l,
		operationLogUseCase: operationLogUseCase,
	}
}

func (s *OperationLogService) SysOperationLogList(ctx context.Context, req *common.SearchListReq) (*pb.SysOperationLogListResp, error) {
	return s.operationLogUseCase.SysOperationLogList(ctx, req)
}
func (s *OperationLogService) SysOperationLogInfo(ctx context.Context, req *pb.SysOperationLogInfoReq) (*pb.SysOperationLogInfoResp, error) {
	return s.operationLogUseCase.SysOperationLogInfo(ctx, req)
}
func (s *OperationLogService) SysOperationLogStore(ctx context.Context, req *pb.SysOperationLogStoreReq) (*pb.SysOperationLogStoreResp, error) {
	return s.operationLogUseCase.SysOperationLogStore(ctx, req)
}

// SysOperationLogStoreProducer 日志生产者
func (s *OperationLogService) SysOperationLogStoreProducer(ctx context.Context, req *pb.SysOperationLogStoreReq) error {
	err := s.operationLogUseCase.SysOperationLogStoreProducer(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

// SysOperationLogStoreConsumer 日志消费者
func (s *OperationLogService) SysOperationLogStoreConsumer(payload []byte) error {
	var params *pb.SysOperationLogStoreReq
	err := json.Unmarshal(payload, params)
	if err != nil {
		return err
	}
	_, err = s.operationLogUseCase.SysOperationLogStore(context.Background(), params)
	if err != nil {
		return err
	}
	return nil
}
