package biz

import (
	"context"
	"fmt"

	pb "fkratos/api/rpc_wechat/v1"
)

// MenuList 公众号-菜单查看
func (o *OfficialAccountUseCase) MenuList(ctx context.Context, req *pb.MenuListReq) (*pb.MenuListReply, error) {
	resp := &pb.MenuListReply{}
	fmt.Println(ctx)
	fmt.Println(req)
	return resp, nil
}
