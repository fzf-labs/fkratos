package biz

import (
	"context"
	v1 "fkratos/api/rpc_sys/v1"

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

func (a *ApiUseCase) SysApiList(ctx context.Context, req *v1.SysApiListReq) (*v1.SysApiListReply, error) {
	resp := new(v1.SysApiListReply)
	list, err := a.sysApiRepo.SysDeptList(ctx, req.GetPermissionId())
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
