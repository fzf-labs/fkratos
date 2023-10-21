package biz

import (
	"context"
	"fmt"

	pb "fkratos/api/rpc_wallet/v1"
)

// WalletDel 钱包-删除多条数据
func (w *WalletUseCase) WalletDel(ctx context.Context, req *pb.WalletDelReq) (*pb.WalletDelReply, error) {
	resp := &pb.WalletDelReply{}
	fmt.Println(ctx)
	fmt.Println(req)
	return resp, nil
}
