package service

import (
	"context"

	pb "fkratos/api/rpc_wallet/v1"
)

func (p *PayService) PayStore(ctx context.Context, req *pb.PayStoreReq) (*pb.PayStoreReply, error) {
	return p.payUseCase.PayStore(ctx, req)
}
