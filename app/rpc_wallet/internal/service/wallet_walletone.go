package service

import (
	"context"

	pb "fkratos/api/rpc_wallet/v1"
)

func (w *WalletService) WalletOne(ctx context.Context, req *pb.WalletOneReq) (*pb.WalletOneReply, error) {
	return w.walletUseCase.WalletOne(ctx, req)
}
