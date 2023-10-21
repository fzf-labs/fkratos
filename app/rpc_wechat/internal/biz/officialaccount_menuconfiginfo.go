package biz

import (
	"context"
	"fmt"

	pb "fkratos/api/rpc_wechat/v1"
)

// MenuConfigInfo 公众号-配置查看
func (o *OfficialAccountUseCase) MenuConfigInfo(ctx context.Context, req *pb.MenuConfigInfoReq) (*pb.MenuConfigInfoReply, error) {
	resp := &pb.MenuConfigInfoReply{}
	fmt.Println(ctx)
	fmt.Println(req)
	return resp, nil
}
