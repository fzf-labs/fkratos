package biz

import (
	"context"
	"fmt"

	pb "fkratos/api/rpc_wechat/v1"
)

// MenuStore 公众号-菜单保存
func (o *OfficialAccountUseCase) MenuStore(ctx context.Context, req *pb.MenuStoreReq) (*pb.MenuStoreReply, error) {
	resp := &pb.MenuStoreReply{}
	fmt.Println(ctx)
	fmt.Println(req)
	return resp, nil
}
