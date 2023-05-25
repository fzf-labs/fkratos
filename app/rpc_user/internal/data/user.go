package data

import (
	"context"
	"fkratos/api/rpc_user/v1"
	"fkratos/app/rpc_user/internal/biz"
	"fkratos/app/rpc_user/internal/data/cache"
	"fkratos/app/rpc_user/internal/data/gorm/userdao"
	"fkratos/app/rpc_user/internal/data/gorm/usermodel"
	"fkratos/internal/errorx"
	"fmt"

	"github.com/fzf-labs/fpkg/util/strutil"
	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.UserRepo = (*UserRepo)(nil)

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	l := log.NewHelper(log.With(logger, "module", "user/repo/user-service"))
	return &UserRepo{
		data: data,
		log:  l,
	}
}

type UserRepo struct {
	data *Data
	log  *log.Helper
}

func (u *UserRepo) CreateUser(ctx context.Context, req *v1.CreateUserReq) (*v1.CreateUserReply, error) {
	reply := new(v1.CreateUserReply)
	return reply, errorx.DataSqlErr
	userDao := userdao.Use(u.data.db).User
	err := userDao.WithContext(ctx).Omit(userDao.ID).Create(&usermodel.User{
		Name: "fuzhifei",
	})
	if err != nil {
		return nil, err
	}
	cacheKey := cache.UUID.BuildCacheKey("fuzhifei")
	rocksCache, err := cacheKey.RocksCache(u.data.rocksCache, ctx, func() (string, error) {
		return strutil.Random(6), nil
	})
	fmt.Println(rocksCache)
	if err != nil {
		return nil, err
	}
	return reply, err
}
func (u *UserRepo) UpdateUser(ctx context.Context, req *v1.UpdateUserReq) (*v1.UpdateUserReply, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepo) DeleteUser(ctx context.Context, req *v1.DeleteUserReq) (*v1.DeleteUserReply, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepo) GetUser(ctx context.Context, req *v1.GetUserReq) (*v1.GetUserReply, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepo) ListUser(ctx context.Context, req *v1.ListUserReq) (*v1.ListUserReply, error) {
	//TODO implement me
	panic("implement me")
}
