package service

import (
	"context"

	pb "fkratos/api/rpc_wallet/v1"
)

// WalletList 钱包-列表数据查询
func (w *WalletService) WalletList(ctx context.Context, req *pb.WalletListReq) (*pb.WalletListReply, error) {
	return w.walletUseCase.WalletList(ctx, req)
}
