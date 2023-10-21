package biz

import (
	"context"
	"fmt"

	pb "fkratos/api/rpc_wallet/v1"
)

// PayDel 支付-删除多条数据
func (p *PayUseCase) PayDel(ctx context.Context, req *pb.PayDelReq) (*pb.PayDelReply, error) {
	resp := &pb.PayDelReply{}
	fmt.Println(ctx)
	fmt.Println(req)
	return resp, nil
}
