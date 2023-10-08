package biz

import (
	"context"
	"encoding/json"
	"fkratos/api/paginator"
	"fkratos/app/rpc_user/internal/data/gorm/fkratos_user_model"
	"fkratos/app/rpc_user/internal/data/gorm/fkratos_user_repo"
	"fkratos/internal/dto"
	"fkratos/internal/errorx"

	pb "fkratos/api/rpc_user/v1"

	"github.com/fzf-labs/fpkg/orm"
	"github.com/go-kratos/kratos/v2/log"
)

type UserGroupRepo interface {
	fkratos_user_repo.IUserGroupRepo
}

type UserGroupUseCase struct {
	log           *log.Helper
	userGroupRepo UserGroupRepo
}

func NewUserGroupUseCase(logger log.Logger, userGroupRepo UserGroupRepo) *UserGroupUseCase {
	l := log.NewHelper(log.With(logger, "module", "biz/userGroup"))
	return &UserGroupUseCase{
		log:           l,
		userGroupRepo: userGroupRepo,
	}
}

func (s *UserGroupUseCase) UserGroupList(ctx context.Context, req *pb.UserGroupListReq) (*pb.UserGroupListReply, error) {
	resp := &pb.UserGroupListReply{
		Paginator: &paginator.PaginatorReply{},
		List:      make([]*pb.UserGroupInfo, 0),
	}
	paginatorReq := &orm.PaginatorReq{}
	err := dto.Copy(paginatorReq, req.GetPaginator())
	if err != nil {
		return nil, errorx.DataFormattingError.WithMetadata(errorx.SetErrMetadata(err))
	}
	result, p, err := s.userGroupRepo.FindMultiByPaginator(ctx, paginatorReq)
	if err != nil {
		return nil, errorx.DataSQLErr.WithMetadata(errorx.SetErrMetadata(err))
	}
	if len(result) > 0 {
		err = dto.Copy(&resp.List, &result)
		if err != nil {
			return nil, errorx.DataFormattingError.WithMetadata(errorx.SetErrMetadata(err))
		}
		err = dto.Copy(&resp.Paginator, p)
		if err != nil {
			return nil, errorx.DataFormattingError.WithMetadata(errorx.SetErrMetadata(err))
		}
	}
	return resp, nil
}

func (s *UserGroupUseCase) UserGroupStore(ctx context.Context, req *pb.UserGroupStoreReq) (*pb.UserGroupStoreReply, error) {
	resp := &pb.UserGroupStoreReply{}
	rolesMarshal, err := json.Marshal(req.GetRoles())
	if err != nil {
		return nil, err
	}
	userGroup := &fkratos_user_model.UserGroup{
		ID:     req.GetId(),
		Name:   req.GetName(),
		Roles:  rolesMarshal,
		Status: int16(req.GetStatus()),
	}
	err = s.userGroupRepo.UpsertOne(ctx, userGroup)
	if err != nil {
		return nil, errorx.DataFormattingError.WithMetadata(errorx.SetErrMetadata(err))
	}
	resp.Id = userGroup.ID
	return resp, nil
}
func (s *UserGroupUseCase) UserGroupDel(ctx context.Context, req *pb.UserGroupDelReq) (*pb.UserGroupDelReply, error) {
	var resp = &pb.UserGroupDelReply{}
	err := s.userGroupRepo.DeleteMultiByIDS(ctx, req.GetIds())
	if err != nil {
		return nil, errorx.DataFormattingError.WithMetadata(errorx.SetErrMetadata(err))
	}
	return resp, nil
}
