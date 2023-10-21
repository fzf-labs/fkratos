package biz

import (
	"context"
	"fmt"

	pb "fkratos/api/rpc_wechat/v1"
)

// MenuConfigStore 公众号-配置保存
func (o *OfficialAccountUseCase) MenuConfigStore(ctx context.Context, req *pb.MenuConfigStoreReq) (*pb.MenuConfigStoreReply, error) {
	resp := &pb.MenuConfigStoreReply{}
	fmt.Println(ctx)
	fmt.Println(req)
	return resp, nil
}
