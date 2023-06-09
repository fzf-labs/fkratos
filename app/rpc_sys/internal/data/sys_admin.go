package data

import (
	"context"
	"fkratos/api/common"
	v1 "fkratos/api/rpc_sys/v1"
	"fkratos/app/rpc_sys/internal/biz"
	"fkratos/app/rpc_sys/internal/data/cache"
	"fkratos/app/rpc_sys/internal/data/gorm/rpc_sys_dao"
	"fkratos/app/rpc_sys/internal/data/gorm/rpc_sys_model"
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
		config: c,
		data:   data,
		log:    l,
	}
}

type SysAdminRepo struct {
	config *conf.Bootstrap
	data   *Data
	log    *log.Helper
}

func (s *SysAdminRepo) GetAdminIdToNameByIds(ctx context.Context, ids []string) (map[string]string, error) {
	res := make(map[string]string)
	sysAdminDao := rpc_sys_dao.Use(s.data.gorm).SysAdmin
	sysAdmins, err := sysAdminDao.WithContext(ctx).Where(sysAdminDao.ID.In(ids...)).Find()
	if err != nil {
		return nil, errorx.DataSqlErr.WithCause(err)
	}
	for _, v := range sysAdmins {
		res[v.ID] = v.Username
	}
	return nil, nil
}

func (s *SysAdminRepo) SysAdminDel(ctx context.Context, ids []string) error {
	sysAdminDao := rpc_sys_dao.Use(s.data.gorm).SysAdmin
	_, err := sysAdminDao.WithContext(ctx).Where(sysAdminDao.ID.In(ids...)).Delete()
	if err != nil {
		return errorx.DataSqlErr.WithCause(err)
	}
	return nil
}

func (s *SysAdminRepo) SysManageStore(ctx context.Context, req *v1.SysManageStoreReq) (*rpc_sys_model.SysAdmin, error) {
	sysAdminDao := rpc_sys_dao.Use(s.data.gorm).SysAdmin
	roleIds, err := jsonutil.Encode(req.RoleIds)
	if err != nil {
		return nil, errorx.DataFormattingError.WithCause(err)
	}
	var sysAdmin *rpc_sys_model.SysAdmin
	if req.Id == "" {
		if req.Avatar == "" {
			req.Avatar = avatar.Url()
		}
		salt := strutil.Random(16)
		pwd, err := crypt.Encrypt(req.Password + salt)
		if err != nil {
			return nil, errorx.DataProcessingError.WithCause(err)
		}
		sysAdmin = &rpc_sys_model.SysAdmin{
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
			return nil, errorx.DataSqlErr.WithCause(err)
		}
	} else {
		sysAdmin, err = sysAdminDao.WithContext(ctx).Where(sysAdminDao.ID.Eq(req.Id)).First()
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, errorx.DataSqlErr.WithCause(err)
		}
		if sysAdmin == nil {
			return nil, errorx.AccountNotExist
		}
		if req.Password != "" {
			pwd, err := crypt.Encrypt(req.Password + sysAdmin.Salt)
			if err != nil {
				return nil, errorx.DataProcessingError.WithCause(err)
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
			return nil, errorx.DataSqlErr.WithCause(err)
		}
	}
	return sysAdmin, nil
}

func (s *SysAdminRepo) SysManageListBySearch(ctx context.Context, req *common.SearchListReq) ([]*rpc_sys_model.SysAdmin, *page.Page, error) {
	sysAdminDao := rpc_sys_dao.Use(s.data.gorm).SysAdmin
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
		return nil, nil, errorx.DataSqlErr.WithCause(err)
	}
	paginator := page.Paginator(int(req.Page), int(req.PageSize), int(total))
	sysAdmins, err := query.Offset(paginator.Offset).Limit(paginator.Limit).Find()
	if err != nil {
		return nil, nil, errorx.DataSqlErr.WithCause(err)
	}
	return sysAdmins, paginator, nil
}

func (s *SysAdminRepo) SysAdminInfoUpdate(ctx context.Context, req *v1.SysAdminInfoUpdateReq) (*v1.SysAdminInfoUpdateReply, error) {
	resp := new(v1.SysAdminInfoUpdateReply)
	sysAdminDao := rpc_sys_dao.Use(s.data.gorm).SysAdmin
	sysAdmin, err := sysAdminDao.WithContext(ctx).Where(sysAdminDao.ID.Eq(req.GetAdminId())).First()
	if err != nil && err != gorm.ErrRecordNotFound {
		return resp, errorx.DataSqlErr.WithCause(err)
	}
	var pwd string
	if req.Password != "" {
		pwd, err = crypt.Encrypt(req.Password + sysAdmin.Salt)
		if err != nil {
			return resp, errorx.DataProcessingError.WithCause(err)
		}
	}
	_, err = sysAdminDao.WithContext(ctx).Where(sysAdminDao.ID.Eq(req.GetAdminId())).Updates(rpc_sys_model.SysAdmin{
		Password: pwd,
		Nickname: req.Nickname,
		Email:    req.Email,
		Mobile:   req.Mobile,
		Motto:    req.Motto,
	})
	if err != nil {
		return nil, errorx.DataSqlErr.WithCause(err)
	}
	cacheKey := cache.SysAdminInfo.NewSingleKey(s.data.redis)
	err = cacheKey.SingleCacheDel(ctx, req.GetAdminId())
	if err != nil {
		return nil, errorx.DataRedisErr.WithCause(err)
	}
	return resp, nil
}

func (s *SysAdminRepo) SysAdminInfoCacheByAdminId(ctx context.Context, adminId string) (*v1.SysAdminInfo, error) {
	resp := new(v1.SysAdminInfo)
	cacheKey := cache.SysAdminInfo.NewSingleKey(s.data.redis)
	res, err := cacheKey.SingleCache(ctx, adminId, func() (string, error) {
		sysAdmin, err := s.SysAdminInfoByAdminId(ctx, adminId)
		if err != nil {
			return "", err
		}
		if sysAdmin == nil {
			return "", nil
		}
		roleIds := make([]string, 0)
		if sysAdmin.RoleIds.String() != "" {
			err = jsonutil.Decode(sysAdmin.RoleIds, &roleIds)
			if err != nil {
				return "", errorx.DataFormattingError.WithCause(err)
			}
		}
		res, err := jsonutil.EncodeToString(v1.SysAdminInfo{
			Id:       sysAdmin.ID,
			Username: sysAdmin.Username,
			Nickname: sysAdmin.Nickname,
			Avatar:   sysAdmin.Avatar,
			Gender:   int32(sysAdmin.Gender),
			Email:    sysAdmin.Email,
			Mobile:   sysAdmin.Mobile,
			JobID:    sysAdmin.JobID,
			DeptID:   sysAdmin.DeptID,
			RoleIds:  roleIds,
			Motto:    sysAdmin.Motto,
		})
		if err != nil {
			return "", errorx.DataFormattingError.WithCause(err)
		}
		return res, nil
	})
	if err != nil {
		return nil, err
	}
	if res != "" {
		err = jsonutil.DecodeString(res, resp)
		if err != nil {
			return nil, errorx.DataFormattingError.WithCause(err)
		}
	}
	return resp, nil
}

// SysAdminInfoByAdminId 根据adminId查询SysAdminInfo
func (s *SysAdminRepo) SysAdminInfoByAdminId(ctx context.Context, adminId string) (*rpc_sys_model.SysAdmin, error) {
	sysAdminDao := rpc_sys_dao.Use(s.data.gorm).SysAdmin
	sysAdmin, err := sysAdminDao.WithContext(ctx).Where(sysAdminDao.ID.Eq(adminId)).First()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errorx.DataSqlErr.WithCause(err)
	}
	return sysAdmin, nil
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
		return nil, errorx.TokenGenerationFailed.WithCause(err)
	}
	err = jwtClient.JwtTokenStore(claims)
	if err != nil {
		return nil, errorx.TokenStorageFailed
	}
	return token, nil
}

// SysAdminInfoByUsername 根据username查询SysAdminInfo
func (s *SysAdminRepo) SysAdminInfoByUsername(ctx context.Context, username string) (*rpc_sys_model.SysAdmin, error) {
	sysAdminDao := rpc_sys_dao.Use(s.data.gorm).SysAdmin
	sysAdmin, err := sysAdminDao.WithContext(ctx).Where(sysAdminDao.Username.Eq(username)).First()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errorx.DataSqlErr.WithCause(err)
	}
	return sysAdmin, nil
}
