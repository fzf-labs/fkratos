package service

import (
	"context"

	pb "fkratos/api/rpc_wechat/v1"
)

// MenuList 公众号-菜单查看
func (o *OfficialAccountService) MenuList(ctx context.Context, req *pb.MenuListReq) (*pb.MenuListReply, error) {
	return o.officialAccountUseCase.MenuList(ctx, req)
}
