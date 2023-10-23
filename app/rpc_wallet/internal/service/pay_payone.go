package service

import (
	"context"

	pb "fkratos/api/rpc_wallet/v1"
)

// PayOne 支付-单条数据查询
func (p *PayService) PayOne(ctx context.Context, req *pb.PayOneReq) (*pb.PayOneReply, error) {
	return p.payUseCase.PayOne(ctx, req)
}
