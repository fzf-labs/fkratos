package biz

import (
	"context"
	sysV1 "fkratos/api/rpc_sys/v1"
	"fkratos/internal/errorx"

	pb "fkratos/api/bff_admin/v1"

	"github.com/jinzhu/copier"
)

// SysPermissionStore 权限-保存
func (s *SysUseCase) SysPermissionStore(ctx context.Context, req *pb.SysPermissionStoreReq) (*pb.SysPermissionStoreResp, error) {
	resp := new(pb.SysPermissionStoreResp)
	result, err := s.sysPermissionClient.SysPermissionStore(ctx, &sysV1.SysPermissionStoreReq{
		Id:        req.GetId(),
		Pid:       req.GetPid(),
		Type:      req.GetType(),
		Title:     req.GetTitle(),
		Name:      req.GetName(),
		Path:      req.GetPath(),
		Icon:      req.GetIcon(),
		MenuType:  req.GetMenuType(),
		Url:       req.GetUrl(),
		Component: req.GetComponent(),
		Keepalive: req.GetKeepalive(),
		Extend:    req.GetExtend(),
		Remark:    req.GetRemark(),
		Sort:      req.GetSort(),
		Status:    req.GetStatus(),
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(resp, result)
	if err != nil {
		return nil, errorx.DataFormattingError.WithError(err).Err()
	}
	return resp, nil
}
