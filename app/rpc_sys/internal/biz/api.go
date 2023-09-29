package biz

import (
	"context"
	v1 "fkratos/api/rpc_sys/v1"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_model"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_repo"

	"github.com/fzf-labs/fpkg/util/timeutil"
	"github.com/go-kratos/kratos/v2/log"
)

type SysAPIRepo interface {
	fkratos_sys_repo.ISysAPIRepo
	GetAPIIdToNameByIds(ctx context.Context, ids []string) (map[string]string, error)
	SysAPIStore(ctx context.Context, model *fkratos_sys_model.SysAPI) (*fkratos_sys_model.SysAPI, error)
}

func NewAPIUseCase(logger log.Logger, sysAPIRepo SysAPIRepo) *APIUseCase {
	l := log.NewHelper(log.With(logger, "module", "rpc_user/biz/api"))
	return &APIUseCase{
		log:        l,
		sysAPIRepo: sysAPIRepo,
	}
}

type APIUseCase struct {
	log        *log.Helper
	sysAPIRepo SysAPIRepo
}

func (a *APIUseCase) SysAPIStore(ctx context.Context, req *v1.SysAPIStoreReq) (*v1.SysAPIStoreReply, error) {
	resp := new(v1.SysAPIStoreReply)
	_, err := a.sysAPIRepo.SysAPIStore(ctx, &fkratos_sys_model.SysAPI{
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

func (a *APIUseCase) SysAPIDel(ctx context.Context, req *v1.SysAPIDelReq) (*v1.SysAPIDelReply, error) {
	resp := new(v1.SysAPIDelReply)
	err := a.sysAPIRepo.DeleteMultiCacheByIDS(ctx, req.GetIds())
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *APIUseCase) SysAPIList(ctx context.Context, req *v1.SysAPIListReq) (*v1.SysAPIListReply, error) {
	resp := new(v1.SysAPIListReply)
	list, err := a.sysAPIRepo.FindMultiByPermissionID(ctx, req.GetPermissionId())
	if err != nil {
		return nil, err
	}
	for _, v := range list {
		resp.List = append(resp.List, &v1.SysAPIInfo{
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
