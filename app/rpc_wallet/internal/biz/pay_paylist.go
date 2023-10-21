package biz

import (
	"context"
	"fmt"

	pb "fkratos/api/rpc_wallet/v1"
)

// PayList 支付-列表数据查询
func (p *PayUseCase) PayList(ctx context.Context, req *pb.PayListReq) (*pb.PayListReply, error) {
	resp := &pb.PayListReply{}
	fmt.Println(ctx)
	fmt.Println(req)
	return resp, nil
}
