package biz

import (
	"context"
	v1 "fkratos/api/rpc_sys/v1"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_model"

	"github.com/fzf-labs/fpkg/util/timeutil"
	"github.com/go-kratos/kratos/v2/log"
)

func NewApiUseCase(logger log.Logger, sysApiRepo SysApiRepo) *ApiUseCase {
	l := log.NewHelper(log.With(logger, "module", "rpc_user/biz/api"))
	return &ApiUseCase{
		log:        l,
		sysApiRepo: sysApiRepo,
	}
}

type ApiUseCase struct {
	log        *log.Helper
	sysApiRepo SysApiRepo
}

func (a *ApiUseCase) SysApiStore(ctx context.Context, req *v1.SysApiStoreReq) (*v1.SysApiStoreReply, error) {
	resp := new(v1.SysApiStoreReply)
	_, err := a.sysApiRepo.SysApiStore(ctx, &fkratos_sys_model.SysAPI{
		PermissionID: req.GetPermissionID(),
		Method:       req.GetMethod(),
		Path:         req.GetPath(),
		Desc:         req.GetDesc(),
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ApiUseCase) SysApiDel(ctx context.Context, req *v1.SysApiDelReq) (*v1.SysApiDelReply, error) {
	resp := new(v1.SysApiDelReply)
	err := a.sysApiRepo.DeleteMultiCacheByIDS(ctx, req.GetIds())
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ApiUseCase) SysApiList(ctx context.Context, req *v1.SysApiListReq) (*v1.SysApiListReply, error) {
	resp := new(v1.SysApiListReply)
	list, err := a.sysApiRepo.FindMultiByPermissionID(ctx, req.GetPermissionId())
	if err != nil {
		return nil, err
	}
	for _, v := range list {
		resp.List = append(resp.List, &v1.SysApiInfo{
			Id:           v.ID,
			PermissionID: v.PermissionID,
			Method:       v.Method,
			Path:         v.Path,
			Desc:         v.Desc,
			CreatedAt:    timeutil.ToDateTimeStringByTime(v.CreatedAt),
			UpdatedAt:    timeutil.ToDateTimeStringByTime(v.UpdatedAt),
		})
	}
	return resp, nil
}
