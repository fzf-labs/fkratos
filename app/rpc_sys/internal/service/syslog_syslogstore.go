package service

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"

	"github.com/fzf-labs/fpkg/util/jsonutil"
)

func (s *SysLogService) SysLogStore(ctx context.Context, req *pb.SysLogStoreReq) (*pb.SysLogStoreResp, error) {
	resp := &pb.SysLogStoreResp{}
	err := s.SysLogStoreProducer(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SysLogStoreProducer 日志生产者
func (s *SysLogService) SysLogStoreProducer(ctx context.Context, req *pb.SysLogStoreReq) error {
	err := s.sysLogUseCase.SysLogStoreProducer(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

// SysLogStoreConsumer 日志消费者
func (s *SysLogService) SysLogStoreConsumer(payload []byte) error {
	var params *pb.SysLogStoreReq
	err := jsonutil.Unmarshal(payload, params)
	if err != nil {
		return err
	}
	_, err = s.sysLogUseCase.SysLogStore(context.Background(), params)
	if err != nil {
		return err
	}
	return nil
}
