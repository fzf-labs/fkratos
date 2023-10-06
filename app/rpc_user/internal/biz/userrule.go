package biz

import (
	"context"
	"fkratos/api/paginator"
	"fkratos/app/rpc_user/internal/data/gorm/fkratos_user_model"
	"fkratos/app/rpc_user/internal/data/gorm/fkratos_user_repo"
	"fkratos/internal/dto"
	"fkratos/internal/errorx"

	pb "fkratos/api/rpc_user/v1"

	"github.com/fzf-labs/fpkg/orm"
	"github.com/go-kratos/kratos/v2/log"
)

type UserRuleRepo interface {
	fkratos_user_repo.IUserRuleRepo
}

type UserRuleUseCase struct {
	log          *log.Helper
	userRuleRepo UserRuleRepo
}

func NewUserRuleUseCase(logger log.Logger, userRuleRepo UserRuleRepo) *UserRuleUseCase {
	l := log.NewHelper(log.With(logger, "module", "biz/userRule"))
	return &UserRuleUseCase{
		log:          l,
		userRuleRepo: userRuleRepo,
	}
}

func (s *UserRuleUseCase) UserRuleList(ctx context.Context, req *pb.UserRuleListReq) (*pb.UserRuleListReply, error) {
	resp := &pb.UserRuleListReply{
		Paginator: &paginator.PaginatorReply{},
		List:      make([]*pb.UserRuleInfo, 0),
	}
	paginatorReq := &orm.PaginatorReq{}
	err := dto.Copy(paginatorReq, req.GetPaginator())
	if err != nil {
		return nil, errorx.DataFormattingError.WithMetadata(errorx.SetErrMetadata(err))
	}
	result, p, err := s.userRuleRepo.FindMultiByPaginator(ctx, paginatorReq)
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
func (s *UserRuleUseCase) UserRuleInfo(ctx context.Context, req *pb.UserRuleInfoReq) (*pb.UserRuleInfoReply, error) {
	resp := &pb.UserRuleInfoReply{
		Info: nil,
	}
	result, err := s.userRuleRepo.FindOneCacheByID(ctx, req.GetID())
	if err != nil {
		return nil, errorx.DataSQLErr.WithMetadata(errorx.SetErrMetadata(err))
	}
	err = dto.Copy(&resp.Info, result)
	if err != nil {
		return nil, errorx.DataFormattingError.WithMetadata(errorx.SetErrMetadata(err))
	}
	return resp, nil
}
func (s *UserRuleUseCase) UserRuleStore(ctx context.Context, req *pb.UserRuleStoreReq) (*pb.UserRuleStoreReply, error) {
	resp := &pb.UserRuleStoreReply{
		ID: "",
	}
	userRule := &fkratos_user_model.UserRule{
		ID:        req.GetId(),
		Pid:       req.GetPid(),
		Type:      req.GetType(),
		Title:     req.GetTitle(),
		Name:      req.GetName(),
		Path:      req.GetPath(),
		Icon:      req.GetIcon(),
		MenuType:  req.GetMenuType(),
		URL:       req.GetUrl(),
		Component: req.GetComponent(),
		Extend:    req.GetExtend(),
		Remark:    req.GetRemark(),
		Status:    int16(req.GetStatus()),
	}
	err := s.userRuleRepo.UpsertOne(ctx, userRule)
	if err != nil {
		return nil, errorx.DataSQLErr.WithMetadata(errorx.SetErrMetadata(err))
	}
	resp.ID = userRule.ID
	return resp, nil
}
func (s *UserRuleUseCase) UserRuleDel(ctx context.Context, req *pb.UserRuleDelReq) (*pb.UserRuleDelReply, error) {
	var resp = &pb.UserRuleDelReply{}
	err := s.userRuleRepo.DeleteMultiCacheByIDS(ctx, req.GetIds())
	if err != nil {
		return nil, errorx.DataSQLErr.WithMetadata(errorx.SetErrMetadata(err))
	}
	return resp, nil
}
