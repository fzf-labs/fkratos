package biz

import (
	"context"
	v1 "fkratos/api/rpc_sys/v1"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_model"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_repo"
	"strings"

	"github.com/fzf-labs/fpkg/util/timeutil"
	"github.com/go-kratos/kratos/v2/log"
)

type SysRoleRepo interface {
	fkratos_sys_repo.ISysRoleRepo
	GetRoleIDToNameByIds(ctx context.Context, ids []string) (map[string]string, error)
	SysRoleList(ctx context.Context) ([]*fkratos_sys_model.SysRole, error)
	SysRoleStore(ctx context.Context, req *v1.SysRoleStoreReq) (*fkratos_sys_model.SysRole, error)
}

func NewRoleUseCase(logger log.Logger, sysRoleRepo SysRoleRepo) *RoleUseCase {
	l := log.NewHelper(log.With(logger, "module", "rpc_user/biz/role"))
	return &RoleUseCase{
		log:         l,
		sysRoleRepo: sysRoleRepo,
	}
}

type RoleUseCase struct {
	log         *log.Helper
	sysRoleRepo SysRoleRepo
}

func (r *RoleUseCase) SysRoleList(ctx context.Context, _ *v1.SysRoleListReq) (*v1.SysRoleListResp, error) {
	resp := new(v1.SysRoleListResp)
	list, err := r.sysRoleRepo.SysRoleList(ctx)
	if err != nil {
		return nil, err
	}
	roles := make([]*v1.SysRoleInfo, 0)
	for _, role := range list {
		roles = append(roles, &v1.SysRoleInfo{
			Id:        role.ID,
			Pid:       role.Pid,
			Name:      role.Name,
			Remark:    role.Remark,
			Status:    int32(role.Status),
			Sort:      int32(role.Sort),
			CreatedAt: role.CreatedAt.Format(timeutil.TimeLayout),
			UpdatedAt: role.UpdatedAt.Format(timeutil.TimeLayout),
			Children:  nil,
		})
	}
	sysRoleGenerateTree(roles)
	return resp, nil
}
func sysRoleGenerateTree(list []*v1.SysRoleInfo) []*v1.SysRoleInfo {
	var trees []*v1.SysRoleInfo
	// Define the top-level root and child nodes
	var roots, childs []*v1.SysRoleInfo
	for _, v := range list {
		if v.Pid != "" {
			// Determine the top-level root node
			roots = append(roots, v)
		}
		childs = append(childs, v)
	}

	for _, v := range roots {
		childTree := &v1.SysRoleInfo{
			Id:        v.Id,
			Pid:       v.Pid,
			Name:      v.Name,
			Remark:    v.Remark,
			Status:    v.Status,
			Sort:      v.Sort,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			Children:  make([]*v1.SysRoleInfo, 0),
		}
		// recursive
		sysRoleRecursiveTree(childTree, childs)

		trees = append(trees, childTree)
	}
	return trees
}
func sysRoleRecursiveTree(tree *v1.SysRoleInfo, allNodes []*v1.SysRoleInfo) {
	for _, v := range allNodes {
		if v.Pid != "" {
			// If the current node is the top-level root node, skip
			continue
		}
		if tree.Id == v.Pid {
			childTree := &v1.SysRoleInfo{
				Id:        v.Id,
				Pid:       v.Pid,
				Name:      v.Name,
				Remark:    v.Remark,
				Status:    v.Status,
				Sort:      v.Sort,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
				Children:  make([]*v1.SysRoleInfo, 0),
			}
			sysRoleRecursiveTree(childTree, allNodes)
			tree.Children = append(tree.Children, childTree)
		}
	}
}

func (r *RoleUseCase) SysRoleInfo(ctx context.Context, req *v1.SysRoleInfoReq) (*v1.SysRoleInfoResp, error) {
	sysRole, err := r.sysRoleRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	split := strings.Split(sysRole.PermissionIds, ",")
	return &v1.SysRoleInfoResp{Info: &v1.SysRoleInfo{
		Id:            sysRole.ID,
		Pid:           sysRole.Pid,
		Name:          sysRole.Name,
		Remark:        sysRole.Remark,
		Status:        int32(sysRole.Status),
		Sort:          int32(sysRole.Sort),
		PermissionIds: split,
		CreatedAt:     timeutil.ToDateTimeStringByTime(sysRole.CreatedAt),
		UpdatedAt:     timeutil.ToDateTimeStringByTime(sysRole.UpdatedAt),
		Children:      nil,
	}}, nil
}

func (r *RoleUseCase) SysRoleStore(ctx context.Context, req *v1.SysRoleStoreReq) (*v1.SysRoleStoreResp, error) {
	_, err := r.sysRoleRepo.SysRoleStore(ctx, req)
	if err != nil {
		return nil, err
	}
	return &v1.SysRoleStoreResp{}, nil
}

func (r *RoleUseCase) SysRoleDel(ctx context.Context, req *v1.SysRoleDelReq) (*v1.SysRoleDelResp, error) {
	err := r.sysRoleRepo.DeleteMultiCacheByIDS(ctx, req.GetIds())
	if err != nil {
		return nil, err
	}
	return &v1.SysRoleDelResp{}, nil
}
