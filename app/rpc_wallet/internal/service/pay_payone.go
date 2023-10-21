package service

import (
	"context"

	pb "fkratos/api/rpc_wallet/v1"
)

func (p *PayService) PayOne(ctx context.Context, req *pb.PayOneReq) (*pb.PayOneReply, error) {
	return p.payUseCase.PayOne(ctx, req)
}
