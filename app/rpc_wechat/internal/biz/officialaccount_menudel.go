package biz

import (
	"context"
	"fmt"

	pb "fkratos/api/rpc_wechat/v1"
)

// MenuDel 公众号-菜单删除
func (o *OfficialAccountUseCase) MenuDel(ctx context.Context, req *pb.MenuDelReq) (*pb.MenuDelReply, error) {
	resp := &pb.MenuDelReply{}
	fmt.Println(ctx)
	fmt.Println(req)
	return resp, nil
}
