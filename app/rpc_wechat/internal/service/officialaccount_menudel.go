package service

import (
	"context"

	pb "fkratos/api/rpc_wechat/v1"
)

// MenuDel 公众号-菜单删除
func (o *OfficialAccountService) MenuDel(ctx context.Context, req *pb.MenuDelReq) (*pb.MenuDelReply, error) {
	return o.officialAccountUseCase.MenuDel(ctx, req)
}
