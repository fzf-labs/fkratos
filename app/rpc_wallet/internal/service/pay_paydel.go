package service

import (
	"context"

	pb "fkratos/api/rpc_wallet/v1"
)

// PayDel 支付-删除多条数据
func (p *PayService) PayDel(ctx context.Context, req *pb.PayDelReq) (*pb.PayDelReply, error) {
	return p.payUseCase.PayDel(ctx, req)
}
