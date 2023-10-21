package biz

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

// SysAuthJwtTokenCheck Auth-Token校验
func (s *SysAuthUseCase) SysAuthJwtTokenCheck(ctx context.Context, req *pb.SysAuthJwtTokenCheckReq) (*pb.SysAuthJwtTokenCheckReply, error) {
	resp := &pb.SysAuthJwtTokenCheckReply{}
	adminID, err := s.sysAdminRepo.SysAuthJwtTokenCheck(ctx, req.GetToken())
	if err != nil {
		return nil, err
	}
	resp.AdminId = adminID
	return resp, nil
}
