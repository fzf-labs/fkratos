package biz

import (
	"context"
	"fmt"

	pb "fkratos/api/rpc_wallet/v1"
)

// WalletStore 钱包-创建一条数据
func (w *WalletUseCase) WalletStore(ctx context.Context, req *pb.WalletStoreReq) (*pb.WalletStoreReply, error) {
	resp := &pb.WalletStoreReply{}
	fmt.Println(ctx)
	fmt.Println(req)
	return resp, nil
}
