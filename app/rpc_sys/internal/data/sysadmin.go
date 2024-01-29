package data

import (
	"context"
	"errors"
	pb "fkratos/api/rpc_sys/v1"
	"fkratos/app/rpc_sys/internal/biz"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_dao"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_model"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_repo"
	"fkratos/internal/errorx"

	conf "github.com/fzf-labs/fkratos-contrib/api/conf/v1"

	"github.com/fzf-labs/fpkg/conv"
	"github.com/fzf-labs/fpkg/crypt"
	"github.com/fzf-labs/fpkg/jwt"
	"github.com/fzf-labs/fpkg/third_api/avatar"
	"github.com/fzf-labs/fpkg/util/jsonutil"
	"github.com/fzf-labs/fpkg/util/strutil"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

var _ biz.SysAdminRepo = (*SysAdminRepo)(nil)

func NewSysAdminRepo(
	logger log.Logger,
	data *Data,
	c *conf.Bootstrap,
	sysAdminRepo *fkratos_sys_repo.SysAdminRepo,
) biz.SysAdminRepo {
	l := log.NewHelper(log.With(logger, "module", "data/sysAdmin"))
	return &SysAdminRepo{
		log:          l,
		data:         data,
		config:       c,
		jwtCache:     NewJwtCache(data.rueidis),
		SysAdminRepo: sysAdminRepo,
	}
}

type SysAdminRepo struct {
	log      *log.Helper
	data     *Data
	config   *conf.Bootstrap
	jwtCache *JwtCache
	*fkratos_sys_repo.SysAdminRepo
}

func (s *SysAdminRepo) GetAdminIDToNameByIds(ctx context.Context, ids []string) (map[string]string, error) {
	res := make(map[string]string)
	sysAdminDao := fkratos_sys_dao.Use(s.data.gorm).SysAdmin
	sysAdmins, err := sysAdminDao.WithContext(ctx).Where(sysAdminDao.ID.In(ids...)).Find()
	if err != nil {
		return nil, errorx.DataSQLErr.WithError(err).Err()
	}
	for _, v := range sysAdmins {
		res[v.ID] = v.Username
	}
	return nil, nil
}

func (s *SysAdminRepo) SysAdminDel(ctx context.Context, ids []string) error {
	sysAdminDao := fkratos_sys_dao.Use(s.data.gorm).SysAdmin
	_, err := sysAdminDao.WithContext(ctx).Where(sysAdminDao.ID.In(ids...)).Delete()
	if err != nil {
		return errorx.DataSQLErr.WithError(err).Err()
	}
	return nil
}

func (s *SysAdminRepo) SysManageStore(ctx context.Context, req *pb.SysManageStoreReq) (*fkratos_sys_model.SysAdmin, error) {
	sysAdminDao := fkratos_sys_dao.Use(s.data.gorm).SysAdmin
	roleIds, err := jsonutil.Marshal(req.RoleIds)
	if err != nil {
		return nil, errorx.DataFormattingError.WithError(err).Err()
	}
	var sysAdmin *fkratos_sys_model.SysAdmin
	if req.Id == "" {
		if req.Avatar == "" {
			req.Avatar = avatar.URL()
		}
		salt := strutil.Random(16)
		pwd, err2 := crypt.Encrypt(req.Password + salt)
		if err2 != nil {
			return nil, errorx.DataProcessingError.WithError(err2).Err()
		}
		sysAdmin = &fkratos_sys_model.SysAdmin{
			Username: req.Username,
			Password: pwd,
			Nickname: req.Nickname,
			Avatar:   req.Avatar,
			Gender:   int16(req.Gender),
			Email:    req.Email,
			Mobile:   req.Mobile,
			JobID:    req.JobID,
			DeptID:   req.DeptID,
			RoleIds:  roleIds,
			Salt:     salt,
			Status:   int16(req.Status),
			Motto:    req.Motto,
		}
		err2 = sysAdminDao.WithContext(ctx).Create(sysAdmin)
		if err2 != nil {
			return nil, errorx.DataSQLErr.WithError(err2).Err()
		}
	} else {
		sysAdmin, err = sysAdminDao.WithContext(ctx).Where(sysAdminDao.ID.Eq(req.Id)).First()
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.DataSQLErr.WithError(err).Err()
		}
		if sysAdmin == nil {
			return nil, errorx.AccountNotExist.Err()
		}
		if req.Password != "" {
			pwd, err2 := crypt.Encrypt(req.Password + sysAdmin.Salt)
			if err2 != nil {
				return nil, errorx.DataProcessingError.WithError(err2).Err()
			}
			sysAdmin.Password = pwd
		}
		_, err = sysAdminDao.WithContext(ctx).Select(
			sysAdminDao.Username,
			sysAdminDao.Nickname,
			sysAdminDao.Avatar,
			sysAdminDao.Gender,
			sysAdminDao.Email,
			sysAdminDao.Mobile,
			sysAdminDao.JobID,
			sysAdminDao.DeptID,
			sysAdminDao.RoleIds,
			sysAdminDao.Status,
			sysAdminDao.Motto,
		).Where(sysAdminDao.ID.Eq(req.Id)).Updates(sysAdmin)
		if err != nil {
			return nil, errorx.DataSQLErr.WithError(err).Err()
		}
	}
	return sysAdmin, nil
}

// ClearJwTToken 清除jwtToken
func (s *SysAdminRepo) ClearJwTToken(_ context.Context, jwtUID string) error {
	jwtClient := jwt.NewJwt(&jwt.Config{
		AccessSecret: s.config.Business.Jwt.AccessSecret,
		AccessExpire: s.config.Business.Jwt.AccessExpire,
		RefreshAfter: s.config.Business.Jwt.RefreshAfter,
		Issuer:       s.config.Business.Jwt.Issuer,
	}, s.jwtCache)
	err := jwtClient.JwtTokenClear(jwtUID)
	if err != nil {
		return err
	}
	return nil
}

// GenerateJwTToken 生成jwtToken
func (s *SysAdminRepo) GenerateJwTToken(_ context.Context, kv map[string]interface{}) (*jwt.Token, error) {
	jwtClient := jwt.NewJwt(&jwt.Config{
		AccessSecret: s.config.Business.Jwt.AccessSecret,
		AccessExpire: s.config.Business.Jwt.AccessExpire,
		RefreshAfter: s.config.Business.Jwt.RefreshAfter,
		Issuer:       s.config.Business.Jwt.Issuer,
	}, s.jwtCache)
	token, claims, err := jwtClient.GenerateToken(kv)
	if err != nil {
		return nil, errorx.TokenGenerationFailed.WithError(err).Err()
	}
	err = jwtClient.JwtTokenStore(claims)
	if err != nil {
		return nil, errorx.TokenStorageFailed.Err()
	}
	return token, nil
}

func (s *SysAdminRepo) SysAuthJwtTokenCheck(_ context.Context, token string) (string, error) {
	///token 解析
	jwtClient := jwt.NewJwt(&jwt.Config{
		AccessSecret: s.config.Business.Jwt.AccessSecret,
		AccessExpire: s.config.Business.Jwt.AccessExpire,
		RefreshAfter: s.config.Business.Jwt.RefreshAfter,
		Issuer:       s.config.Business.Jwt.Issuer,
	}, s.jwtCache)
	claims, err := jwtClient.ParseToken(token)
	if err != nil {
		if errors.Is(err, jwt.TokenExpired) {
			return "", errorx.TokenExpired.Err()
		}
		return "", errorx.TokenInvalidErr.Err()
	}
	return conv.String(claims["uid"]), nil
}
