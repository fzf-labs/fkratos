package biz

import (
	"context"
	"fkratos/app/rpc_user/internal/data/gorm/fkratos_user_model"
	"fkratos/internal/errorx"

	pb "fkratos/api/rpc_user/v1"

	"github.com/fzf-labs/fpkg/util/jsonutil"
)

// UserGroupStore 用户分组-保存
func (u *UserGroupUseCase) UserGroupStore(ctx context.Context, req *pb.UserGroupStoreReq) (*pb.UserGroupStoreReply, error) {
	resp := &pb.UserGroupStoreReply{}
	rolesMarshal, err := jsonutil.Marshal(req.GetRoles())
	if err != nil {
		return nil, err
	}
	userGroup := &fkratos_user_model.UserGroup{
		ID:     req.GetId(),
		Name:   req.GetName(),
		Roles:  rolesMarshal,
		Status: int16(req.GetStatus()),
	}
	err = u.userGroupRepo.UpsertOne(ctx, userGroup)
	if err != nil {
		return nil, errorx.DataFormattingError.WithError(err).Err()
	}
	resp.Id = userGroup.ID
	return resp, nil
}
