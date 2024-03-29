package biz

import (
	"context"
	"encoding/json"
	pb "fkratos/api/rpc_sys/v1"
)

// SysLogStore 日志-保存
func (s *SysLogUseCase) SysLogStore(ctx context.Context, req *pb.SysLogStoreReq) (*pb.SysLogStoreResp, error) {
	resp := &pb.SysLogStoreResp{}
	_, err := s.sysLogRepo.SysLogStore(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SysLogUseCase) SysLogStoreProducer(ctx context.Context, req *pb.SysLogStoreReq) error {
	err := s.sysLogRepo.SysLogStoreMQProducer(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func (s *SysLogUseCase) SysLogStoreConsumer(payload []byte) error {
	req := &pb.SysLogStoreReq{}
	err := json.Unmarshal(payload, req)
	if err != nil {
		return err
	}
	_, err = s.sysLogRepo.SysLogStore(context.Background(), req)
	if err != nil {
		return err
	}
	return nil
}
