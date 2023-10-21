package biz

import (
	"context"
	"fmt"

	pb "fkratos/api/rpc_wallet/v1"
)

// WalletOne 钱包-单条数据查询
func (w *WalletUseCase) WalletOne(ctx context.Context, req *pb.WalletOneReq) (*pb.WalletOneReply, error) {
	resp := &pb.WalletOneReply{}
	fmt.Println(ctx)
	fmt.Println(req)
	return resp, nil
}
