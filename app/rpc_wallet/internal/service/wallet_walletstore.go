package service

import (
	"context"

	pb "fkratos/api/rpc_wallet/v1"
)

// WalletStore 钱包-创建一条数据
func (w *WalletService) WalletStore(ctx context.Context, req *pb.WalletStoreReq) (*pb.WalletStoreReply, error) {
	return w.walletUseCase.WalletStore(ctx, req)
}
