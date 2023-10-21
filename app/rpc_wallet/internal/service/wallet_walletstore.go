package service

import (
	"context"

	pb "fkratos/api/rpc_wallet/v1"
)

func (w *WalletService) WalletStore(ctx context.Context, req *pb.WalletStoreReq) (*pb.WalletStoreReply, error) {
	return w.walletUseCase.WalletStore(ctx, req)
}
