package data

import (
	"context"
	"errors"
	v1 "fkratos/api/rpc_sys/v1"
	"fkratos/app/rpc_sys/internal/biz"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_dao"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_model"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_repo"
	"fkratos/internal/bootstrap/conf"
	"fkratos/internal/errorx"

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
	c *conf.Bootstrap,
	data *Data,
	logger log.Logger,
	sysAdminRepo *fkratos_sys_repo.SysAdminRepo,
) biz.SysAdminRepo {
	l := log.NewHelper(log.With(logger, "module", "rpc_sys/data/sys_admin"))
	return &SysAdminRepo{
		config:       c,
		log:          l,
		data:         data,
		jwtCache:     NewJwtCache(data.rueidis),
		SysAdminRepo: sysAdminRepo,
	}
}

type SysAdminRepo struct {
	config   *conf.Bootstrap
	log      *log.Helper
	data     *Data
	jwtCache *JwtCache
	*fkratos_sys_repo.SysAdminRepo
}

func (s *SysAdminRepo) GetAdminIDToNameByIds(ctx context.Context, ids []string) (map[string]string, error) {
	res := make(map[string]string)
	sysAdminDao := fkratos_sys_dao.Use(s.data.gorm).SysAdmin
	sysAdmins, err := sysAdminDao.WithContext(ctx).Where(sysAdminDao.ID.In(ids...)).Find()
	if err != nil {
		return nil, errorx.DataSQLErr.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
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
		return errorx.DataSQLErr.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return nil
}

func (s *SysAdminRepo) SysManageStore(ctx context.Context, req *v1.SysManageStoreReq) (*fkratos_sys_model.SysAdmin, error) {
	sysAdminDao := fkratos_sys_dao.Use(s.data.gorm).SysAdmin
	roleIds, err := jsonutil.Encode(req.RoleIds)
	if err != nil {
		return nil, errorx.DataFormattingError.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	var sysAdmin *fkratos_sys_model.SysAdmin
	if req.Id == "" {
		if req.Avatar == "" {
			req.Avatar = avatar.URL()
		}
		salt := strutil.Random(16)
		pwd, err2 := crypt.Encrypt(req.Password + salt)
		if err2 != nil {
			return nil, errorx.DataProcessingError.WithCause(err2).WithMetadata(errorx.SetErrMetadata(err2))
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
			return nil, errorx.DataSQLErr.WithCause(err2).WithMetadata(errorx.SetErrMetadata(err2))
		}
	} else {
		sysAdmin, err = sysAdminDao.WithContext(ctx).Where(sysAdminDao.ID.Eq(req.Id)).First()
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.DataSQLErr.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
		}
		if sysAdmin == nil {
			return nil, errorx.AccountNotExist
		}
		if req.Password != "" {
			pwd, err2 := crypt.Encrypt(req.Password + sysAdmin.Salt)
			if err2 != nil {
				return nil, errorx.DataProcessingError.WithCause(err2).WithMetadata(errorx.SetErrMetadata(err2))
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
			return nil, errorx.DataSQLErr.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
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
		return nil, errorx.TokenGenerationFailed.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	err = jwtClient.JwtTokenStore(claims)
	if err != nil {
		return nil, errorx.TokenStorageFailed
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
			return "", errorx.TokenExpired
		}
		return "", errorx.TokenInvalidErr
	}
	return conv.String(claims["uid"]), nil
}
