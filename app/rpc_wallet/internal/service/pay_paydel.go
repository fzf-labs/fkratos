package service

import (
	"context"

	pb "fkratos/api/rpc_wallet/v1"
)

func (p *PayService) PayDel(ctx context.Context, req *pb.PayDelReq) (*pb.PayDelReply, error) {
	return p.payUseCase.PayDel(ctx, req)
}
