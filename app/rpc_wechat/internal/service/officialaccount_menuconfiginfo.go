package service

import (
	"context"

	pb "fkratos/api/rpc_wechat/v1"
)

func (o *OfficialAccountService) MenuConfigInfo(ctx context.Context, req *pb.MenuConfigInfoReq) (*pb.MenuConfigInfoReply, error) {
	return o.officialAccountUseCase.MenuConfigInfo(ctx, req)
}
