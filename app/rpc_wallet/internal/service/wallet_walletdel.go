package service

import (
	"context"

	pb "fkratos/api/rpc_wallet/v1"
)

// WalletDel 钱包-删除多条数据
func (w *WalletService) WalletDel(ctx context.Context, req *pb.WalletDelReq) (*pb.WalletDelReply, error) {
	return w.walletUseCase.WalletDel(ctx, req)
}
