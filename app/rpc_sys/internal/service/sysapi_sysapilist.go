package service

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

func (s *SysAPIService) SysAPIList(ctx context.Context, req *pb.SysAPIListReq) (*pb.SysAPIListReply, error) {
	return s.sysAPIUseCase.SysAPIList(ctx, req)
}
