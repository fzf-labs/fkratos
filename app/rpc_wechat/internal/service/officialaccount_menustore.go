package service

import (
	"context"

	pb "fkratos/api/rpc_wechat/v1"
)

func (o *OfficialAccountService) MenuStore(ctx context.Context, req *pb.MenuStoreReq) (*pb.MenuStoreReply, error) {
	return o.officialAccountUseCase.MenuStore(ctx, req)
}
