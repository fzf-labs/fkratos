package biz

import (
	"context"
	v1 "fkratos/api/rpc_sys/v1"

	"github.com/fzf-labs/fpkg/util/timeutil"
	"github.com/go-kratos/kratos/v2/log"
)

func NewPermissionUseCase(logger log.Logger, sysPermissionRepo SysPermissionRepo) *PermissionUseCase {
	l := log.NewHelper(log.With(logger, "module", "rpc_user/biz/permission"))
	return &PermissionUseCase{
		log:               l,
		sysPermissionRepo: sysPermissionRepo,
	}
}

type PermissionUseCase struct {
	log               *log.Helper
	sysPermissionRepo SysPermissionRepo
}

func (p *PermissionUseCase) SysPermissionList(ctx context.Context, req *v1.SysPermissionListReq) (*v1.SysPermissionListResp, error) {
	resp := new(v1.SysPermissionListResp)
	list, err := p.sysPermissionRepo.SysPermissionList(ctx)
	if err != nil {
		return nil, err
	}
	roles := make([]*v1.SysPermissionInfo, 0)
	for _, v := range list {
		roles = append(roles, &v1.SysPermissionInfo{
			Id:        v.ID,
			Pid:       v.Pid,
			Type:      v.Type,
			Title:     v.Title,
			Name:      v.Name,
			Path:      v.Path,
			Icon:      v.Icon,
			MenuType:  v.MenuType,
			Url:       v.URL,
			Component: v.Component,
			Extend:    v.Extend,
			Remark:    v.Remark,
			Sort:      int32(v.Sort),
			Status:    int32(v.Status),
			CreatedAt: timeutil.ToDateTimeStringByTime(v.CreatedAt.Time),
			UpdatedAt: timeutil.ToDateTimeStringByTime(v.UpdatedAt.Time),
			Children:  nil,
		})
	}
	resp.List = sysPermMenuGenerateTree(roles)
	return resp, nil
}
func sysPermMenuGenerateTree(list []*v1.SysPermissionInfo) []*v1.SysPermissionInfo {
	var trees []*v1.SysPermissionInfo
	// Define the top-level root and child nodes
	var roots, childs []*v1.SysPermissionInfo
	for _, v := range list {
		if v.Pid != "" {
			// Determine the top-level root node
			roots = append(roots, v)
		}
		childs = append(childs, v)
	}

	for _, v := range roots {
		childTree := &v1.SysPermissionInfo{
			Id:        v.Id,
			Pid:       v.Pid,
			Type:      v.Type,
			Title:     v.Title,
			Name:      v.Name,
			Path:      v.Path,
			Icon:      v.Icon,
			MenuType:  v.MenuType,
			Url:       v.Url,
			Component: v.Component,
			Keepalive: v.Keepalive,
			Extend:    v.Extend,
			Remark:    v.Remark,
			Sort:      v.Sort,
			Status:    v.Status,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			Children:  make([]*v1.SysPermissionInfo, 0),
		}
		// recursive
		sysPermMenuRecursiveTree(childTree, childs)

		trees = append(trees, childTree)
	}
	return trees
}
func sysPermMenuRecursiveTree(tree *v1.SysPermissionInfo, allNodes []*v1.SysPermissionInfo) {
	for _, v := range allNodes {
		if v.Pid == "" {
			// If the current node is the top-level root node, skip
			continue
		}
		if tree.Id == v.Pid {
			childTree := &v1.SysPermissionInfo{
				Id:        v.Id,
				Pid:       v.Pid,
				Type:      v.Type,
				Title:     v.Title,
				Name:      v.Name,
				Path:      v.Path,
				Icon:      v.Icon,
				MenuType:  v.MenuType,
				Url:       v.Url,
				Component: v.Component,
				Keepalive: v.Keepalive,
				Extend:    v.Extend,
				Remark:    v.Remark,
				Sort:      v.Sort,
				Status:    v.Status,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
				Children:  make([]*v1.SysPermissionInfo, 0),
			}
			sysPermMenuRecursiveTree(childTree, allNodes)
			tree.Children = append(tree.Children, childTree)
		}
	}
}

func (p *PermissionUseCase) SysPermissionInfo(ctx context.Context, req *v1.SysPermissionInfoReq) (*v1.SysPermissionInfoResp, error) {
	_, err := p.sysPermissionRepo.SysPermissionInfoById(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &v1.SysPermissionInfoResp{}, nil
}

func (p *PermissionUseCase) SysPermissionStore(ctx context.Context, req *v1.SysPermissionStoreReq) (*v1.SysPermissionStoreResp, error) {
	_, err := p.sysPermissionRepo.SysPermissionStore(ctx, req)
	if err != nil {
		return nil, err
	}
	return &v1.SysPermissionStoreResp{}, nil
}

func (p *PermissionUseCase) SysPermissionDel(ctx context.Context, req *v1.SysPermissionDelReq) (*v1.SysPermissionDelResp, error) {
	err := p.sysPermissionRepo.SysPermissionDelByIds(ctx, req.GetIds())
	if err != nil {
		return nil, err
	}
	return &v1.SysPermissionDelResp{}, nil
}

func (p *PermissionUseCase) SysPermissionStatus(ctx context.Context, req *v1.SysPermissionStatusReq) (*v1.SysPermissionStatusResp, error) {
	err := p.sysPermissionRepo.SysPermissionUpdateStatus(ctx, req.GetId(), req.GetStatus())
	if err != nil {
		return nil, err
	}
	return &v1.SysPermissionStatusResp{}, nil
}
