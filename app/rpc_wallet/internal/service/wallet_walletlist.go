package service

import (
	"context"

	pb "fkratos/api/rpc_wallet/v1"
)

func (w *WalletService) WalletList(ctx context.Context, req *pb.WalletListReq) (*pb.WalletListReply, error) {
	return w.walletUseCase.WalletList(ctx, req)
}
