package biz

import (
	"context"
	"fkratos/app/rpc_user/internal/data/gorm/fkratos_user_model"
	"fkratos/internal/errorx"

	pb "fkratos/api/rpc_user/v1"

	"github.com/fzf-labs/fpkg/crypt"
	"github.com/fzf-labs/fpkg/util/uuidutil"
)

// UserStore 用户-保存
func (u *UserUseCase) UserStore(ctx context.Context, req *pb.UserStoreReq) (*pb.UserStoreReply, error) {
	if req.Id != "" {
		userInfo, err2 := u.userRepo.FindOneCacheByID(ctx, req.Id)
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
		err := u.userRepo.UpdateOne(ctx, user)
		if err != nil {
			return nil, errorx.DataSQLErr.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
		}
		return &pb.UserStoreReply{
			Id: user.ID,
		}, nil
	} else {
		findOneCacheByUsername, err := u.userRepo.FindOneCacheByUsername(ctx, req.GetUsername())
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
		err = u.userRepo.CreateOne(ctx, user)
		if err != nil {
			return nil, errorx.DataSQLErr.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
		}
		return &pb.UserStoreReply{
			Id: user.ID,
		}, nil
	}
}
