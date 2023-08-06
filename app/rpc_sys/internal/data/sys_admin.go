package data

import (
	"context"
	"errors"
	"fkratos/api/paginator"
	v1 "fkratos/api/rpc_sys/v1"
	"fkratos/app/rpc_sys/internal/biz"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_dao"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_model"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_repo"
	"fkratos/internal/bootstrap/conf"
	"fkratos/internal/errorx"
	"strings"

	"github.com/fzf-labs/fpkg/conv"
	"github.com/fzf-labs/fpkg/crypt"
	"github.com/fzf-labs/fpkg/jwt"
	"github.com/fzf-labs/fpkg/page"
	"github.com/fzf-labs/fpkg/third_api/avatar"
	"github.com/fzf-labs/fpkg/util/jsonutil"
	"github.com/fzf-labs/fpkg/util/strutil"
	"github.com/fzf-labs/fpkg/util/timeutil"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

var _ biz.SysAdminRepo = (*SysAdminRepo)(nil)

func NewSysAdminRepo(c *conf.Bootstrap, data *Data, logger log.Logger) biz.SysAdminRepo {
	l := log.NewHelper(log.With(logger, "module", "rpc_sys/data/sys_admin"))
	return &SysAdminRepo{
		config:       c,
		data:         data,
		log:          l,
		SysAdminRepo: fkratos_sys_repo.NewSysAdminRepo(data.gorm, data.redis),
	}
}

type SysAdminRepo struct {
	config *conf.Bootstrap
	log    *log.Helper
	data   *Data
	*fkratos_sys_repo.SysAdminRepo
}

func (s *SysAdminRepo) GetAdminIdToNameByIds(ctx context.Context, ids []string) (map[string]string, error) {
	res := make(map[string]string)
	sysAdminDao := fkratos_sys_dao.Use(s.data.gorm).SysAdmin
	sysAdmins, err := sysAdminDao.WithContext(ctx).Where(sysAdminDao.ID.In(ids...)).Find()
	if err != nil {
		return nil, errorx.DataSqlErr.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
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
		return errorx.DataSqlErr.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
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
			req.Avatar = avatar.Url()
		}
		salt := strutil.Random(16)
		pwd, err := crypt.Encrypt(req.Password + salt)
		if err != nil {
			return nil, errorx.DataProcessingError.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
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
		err = sysAdminDao.WithContext(ctx).Create(sysAdmin)
		if err != nil {
			return nil, errorx.DataSqlErr.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
		}
	} else {
		sysAdmin, err = sysAdminDao.WithContext(ctx).Where(sysAdminDao.ID.Eq(req.Id)).First()
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.DataSqlErr.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
		}
		if sysAdmin == nil {
			return nil, errorx.AccountNotExist
		}
		if req.Password != "" {
			pwd, err := crypt.Encrypt(req.Password + sysAdmin.Salt)
			if err != nil {
				return nil, errorx.DataProcessingError.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
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
			return nil, errorx.DataSqlErr.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
		}
	}
	return sysAdmin, nil
}

func (s *SysAdminRepo) SysManageListBySearch(ctx context.Context, req *paginator.PaginatorReq) ([]*fkratos_sys_model.SysAdmin, *page.Page, error) {
	sysAdminDao := fkratos_sys_dao.Use(s.data.gorm).SysAdmin
	query := sysAdminDao.WithContext(ctx)
	if req.QuickSearch != "" {
		query = query.Where(sysAdminDao.Nickname.Like(req.QuickSearch))
	} else {
		for _, search := range req.Search {
			if search.Field == "id" {
				query = query.Where(sysAdminDao.ID.Eq(search.Val))
			}
			if search.Field == "username" {
				query = query.Where(sysAdminDao.Username.Eq(search.Val))
			}
			if search.Field == "nickname" {
				query = query.Where(sysAdminDao.Nickname.Eq(search.Val))
			}
			if search.Field == "mobile" {
				query = query.Where(sysAdminDao.Mobile.Eq(search.Val))
			}
			if search.Field == "email" {
				query = query.Where(sysAdminDao.Email.Eq(search.Val))
			}
			if search.Field == "status" {
				query = query.Where(sysAdminDao.Status.Eq(conv.Int16(search.Val)))
			}
			if search.Field == "createdAt" {
				ss := strings.Split(search.Val, ",")
				query = query.Where(sysAdminDao.CreatedAt.Gte(timeutil.Carbon().Parse(ss[0]).Carbon2Time()), sysAdminDao.CreatedAt.Lte(timeutil.Carbon().Parse(ss[1]).Carbon2Time()))
			}
		}
	}

	queryCount := query
	total, err := queryCount.Count()
	if err != nil {
		return nil, nil, errorx.DataSqlErr.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	paginator := page.Paginator(int(req.Page), int(req.PageSize), int(total))
	sysAdmins, err := query.Offset(paginator.Offset).Limit(paginator.Limit).Find()
	if err != nil {
		return nil, nil, errorx.DataSqlErr.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return sysAdmins, paginator, nil
}

// ClearJwTToken 清除jwtToken
func (s *SysAdminRepo) ClearJwTToken(ctx context.Context, jwtUId string) error {
	jwtClient := jwt.NewJwt(s.data.redis, &jwt.Config{
		AccessSecret: s.config.Business.Jwt.AccessSecret,
		AccessExpire: s.config.Business.Jwt.AccessExpire,
		RefreshAfter: s.config.Business.Jwt.RefreshAfter,
		Issuer:       s.config.Business.Jwt.Issuer,
	})
	err := jwtClient.JwtTokenClear(jwtUId)
	if err != nil {
		return err
	}
	return nil
}

// GenerateJwTToken 生成jwtToken
func (s *SysAdminRepo) GenerateJwTToken(ctx context.Context, kv map[string]interface{}) (*jwt.Token, error) {
	jwtClient := jwt.NewJwt(s.data.redis, &jwt.Config{
		AccessSecret: s.config.Business.Jwt.AccessSecret,
		AccessExpire: s.config.Business.Jwt.AccessExpire,
		RefreshAfter: s.config.Business.Jwt.RefreshAfter,
		Issuer:       s.config.Business.Jwt.Issuer,
	})
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

func (s *SysAdminRepo) SysAuthJwtTokenCheck(ctx context.Context, token string) (string, error) {
	///token 解析
	jwtClient := jwt.NewJwt(s.data.redis, &jwt.Config{
		AccessSecret: s.config.Business.Jwt.AccessSecret,
		AccessExpire: s.config.Business.Jwt.AccessExpire,
		RefreshAfter: s.config.Business.Jwt.RefreshAfter,
		Issuer:       s.config.Business.Jwt.Issuer,
	})
	claims, err := jwtClient.ParseToken(token)
	if err != nil {
		if errors.Is(err, jwt.TokenExpired) {
			return "", errorx.TokenExpired
		}
		return "", errorx.TokenInvalidErr
	}
	return conv.String(claims["adminId"]), nil
}
