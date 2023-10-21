package service

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

func (s *SysAPIService) SysAPIStore(ctx context.Context, req *pb.SysAPIStoreReq) (*pb.SysAPIStoreReply, error) {
	return s.sysAPIUseCase.SysAPIStore(ctx, req)
}
