package service

import (
	"context"

	pb "fkratos/api/rpc_wechat/v1"
)

// MenuConfigStore 公众号-配置保存
func (o *OfficialAccountService) MenuConfigStore(ctx context.Context, req *pb.MenuConfigStoreReq) (*pb.MenuConfigStoreReply, error) {
	return o.officialAccountUseCase.MenuConfigStore(ctx, req)
}
