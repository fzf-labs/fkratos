package service

import (
	"context"

	pb "fkratos/api/rpc_wechat/v1"
)

func (o *OfficialAccountService) MenuList(ctx context.Context, req *pb.MenuListReq) (*pb.MenuListReply, error) {
	return o.officialAccountUseCase.MenuList(ctx, req)
}
