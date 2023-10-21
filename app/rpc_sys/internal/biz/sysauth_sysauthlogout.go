package biz

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

// SysAuthLogout Auth-退出
func (s *SysAuthUseCase) SysAuthLogout(ctx context.Context, req *pb.SysAuthLogoutReq) (*pb.SysAuthLogoutReply, error) {
	resp := &pb.SysAuthLogoutReply{}
	err := s.sysAdminRepo.ClearJwTToken(ctx, req.GetAdminId())
	if err != nil {
		return nil, err
	}
	return resp, nil
}
