package biz

import (
	"context"
	"fkratos/api/paginator"
	pb "fkratos/api/rpc_user/v1"
	"fkratos/app/rpc_user/internal/data/gorm/fkratos_user_model"
	"fkratos/app/rpc_user/internal/data/gorm/fkratos_user_repo"
	"fkratos/internal/dto"
	"fkratos/internal/errorx"

	"github.com/fzf-labs/fpkg/crypt"
	"github.com/fzf-labs/fpkg/orm"
	"github.com/fzf-labs/fpkg/util/uuidutil"
	"github.com/go-kratos/kratos/v2/log"
)

type UserRepo interface {
	fkratos_user_repo.IUserRepo
}

type UserUseCase struct {
	log      *log.Helper
	userRepo UserRepo
}

func NewUserUseCase(logger log.Logger, userRepo UserRepo) *UserUseCase {
	l := log.NewHelper(log.With(logger, "module", "biz/user"))
	return &UserUseCase{
		log:      l,
		userRepo: userRepo,
	}
}

func (s *UserUseCase) UserList(ctx context.Context, req *pb.UserListReq) (*pb.UserListReply, error) {
	resp := &pb.UserListReply{
		Paginator: &paginator.PaginatorReply{},
		List:      make([]*pb.UserInfo, 0),
	}
	paginatorReq := &orm.PaginatorReq{}
	err := dto.Copy(paginatorReq, req.GetPaginator())
	if err != nil {
		return nil, errorx.DataFormattingError.WithMetadata(errorx.SetErrMetadata(err))
	}
	result, p, err := s.userRepo.FindMultiByPaginator(ctx, paginatorReq)
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
func (s *UserUseCase) UserInfo(ctx context.Context, req *pb.UserInfoReq) (*pb.UserInfoReply, error) {
	resp := &pb.UserInfoReply{
		Info: &pb.UserInfo{},
	}
	result, err := s.userRepo.FindOneCacheByUID(ctx, req.GetUid())
	if err != nil {
		return nil, err
	}
	err = dto.Copy(&resp.Info, result)
	if err != nil {
		return nil, errorx.DataFormattingError.WithMetadata(errorx.SetErrMetadata(err))
	}
	return resp, nil
}

func (s *UserUseCase) UserStore(ctx context.Context, req *pb.UserStoreReq) (*pb.UserStoreReply, error) {
	if req.Id != "" {
		userInfo, err2 := s.userRepo.FindOneCacheByID(ctx, req.Id)
		if err2 != nil {
			return nil, err2
		}
		user := &fkratos_user_model.User{
			ID:          req.GetId(),
			UID:         userInfo.UID,
			UserGroupID: req.GetUserGroupID(),
			Username:    req.GetUsername(),
			Phone:       req.GetPhone(),
			Email:       req.GetEmail(),
			Password:    userInfo.Password,
			Salt:        userInfo.Salt,
			Nickname:    req.GetNickname(),
			Sex:         int16(req.GetSex()),
			Avatar:      req.GetAvatar(),
			Profile:     req.GetProfile(),
			Status:      int16(req.GetStatus()),
		}
		if req.GetPassword() != "" {
			encrypt, err := crypt.Encrypt(req.GetPassword() + userInfo.Salt)
			if err != nil {
				return nil, err
			}
			user.Password = encrypt
		}
		err := s.userRepo.UpdateOne(ctx, user)
		if err != nil {
			return nil, errorx.DataSQLErr.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
		}
		return &pb.UserStoreReply{
			Id: user.ID,
		}, nil
	} else {
		findOneCacheByUsername, err := s.userRepo.FindOneCacheByUsername(ctx, req.GetUsername())
		if err != nil {
			return nil, err
		}
		if findOneCacheByUsername != nil {
			return nil, errorx.AccountDuplicateUsername
		}
		salt := uuidutil.KSUID()
		user := &fkratos_user_model.User{
			ID:          req.GetId(),
			UID:         uuidutil.KSUIDByTime(),
			UserGroupID: req.GetUserGroupID(),
			Username:    req.GetUsername(),
			Phone:       req.GetPhone(),
			Email:       req.GetEmail(),
			Password:    req.GetPassword(),
			Salt:        salt,
			Nickname:    req.GetNickname(),
			Sex:         int16(req.GetSex()),
			Avatar:      req.GetAvatar(),
			Profile:     req.GetProfile(),
			Other:       nil,
			Status:      int16(req.GetStatus()),
		}
		encrypt, err := crypt.Encrypt(req.GetPassword() + salt)
		if err != nil {
			return nil, err
		}
		user.Password = encrypt
		err = s.userRepo.CreateOne(ctx, user)
		if err != nil {
			return nil, errorx.DataSQLErr.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
		}
		return &pb.UserStoreReply{
			Id: user.ID,
		}, nil
	}
}

func (s *UserUseCase) UserDel(ctx context.Context, req *pb.UserDelReq) (*pb.UserDelReply, error) {
	resp := new(pb.UserDelReply)
	err := s.userRepo.DeleteMultiCacheByIDS(ctx, req.GetIds())
	if err != nil {
		return nil, err
	}
	return resp, nil
}
