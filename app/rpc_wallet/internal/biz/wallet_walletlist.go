package biz

import (
	"context"
	"fmt"

	pb "fkratos/api/rpc_wallet/v1"
)

// WalletList 钱包-列表数据查询
func (w *WalletUseCase) WalletList(ctx context.Context, req *pb.WalletListReq) (*pb.WalletListReply, error) {
	resp := &pb.WalletListReply{}
	fmt.Println(ctx)
	fmt.Println(req)
	return resp, nil
}
