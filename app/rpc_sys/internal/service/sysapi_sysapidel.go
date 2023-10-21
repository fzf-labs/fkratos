package service

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

func (s *SysAPIService) SysAPIDel(ctx context.Context, req *pb.SysAPIDelReq) (*pb.SysAPIDelReply, error) {
	return s.sysAPIUseCase.SysAPIDel(ctx, req)
}
