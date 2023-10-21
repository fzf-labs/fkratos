package biz

import (
	"context"
	"fmt"

	pb "fkratos/api/rpc_wallet/v1"
)

// PayStore 支付-创建一条数据
func (p *PayUseCase) PayStore(ctx context.Context, req *pb.PayStoreReq) (*pb.PayStoreReply, error) {
	resp := &pb.PayStoreReply{}
	fmt.Println(ctx)
	fmt.Println(req)
	return resp, nil
}
