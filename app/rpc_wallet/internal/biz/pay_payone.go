package biz

import (
	"context"
	"fmt"

	pb "fkratos/api/rpc_wallet/v1"
)

// PayOne 支付-单条数据查询
func (p *PayUseCase) PayOne(ctx context.Context, req *pb.PayOneReq) (*pb.PayOneReply, error) {
	resp := &pb.PayOneReply{}
	fmt.Println(ctx)
	fmt.Println(req)
	return resp, nil
}
