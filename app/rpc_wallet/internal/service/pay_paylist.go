package service

import (
	"context"

	pb "fkratos/api/rpc_wallet/v1"
)

func (p *PayService) PayList(ctx context.Context, req *pb.PayListReq) (*pb.PayListReply, error) {
	return p.payUseCase.PayList(ctx, req)
}
