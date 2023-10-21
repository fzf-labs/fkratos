package biz

import (
	"context"
	pb "fkratos/api/rpc_sys/v1"

	"github.com/fzf-labs/fpkg/util/timeutil"
)

// SysDeptInfo 部门-单个部门信息
func (s *SysDeptUseCase) SysDeptInfo(ctx context.Context, req *pb.SysDeptInfoReq) (*pb.SysDeptInfoReply, error) {
	resp := &pb.SysDeptInfoReply{}
	sysDept, err := s.sysDeptRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	resp.Info = &pb.SysDeptInfo{
		Id:          sysDept.ID,
		Pid:         sysDept.Pid,
		Name:        sysDept.Name,
		FullName:    sysDept.FullName,
		Responsible: sysDept.Responsible,
		Phone:       sysDept.Phone,
		Email:       sysDept.Email,
		Type:        int32(sysDept.Type),
		Status:      int32(sysDept.Status),
		Sort:        int32(sysDept.Sort),
		CreatedAt:   timeutil.ToDateTimeStringByTime(sysDept.CreatedAt),
		UpdatedAt:   timeutil.ToDateTimeStringByTime(sysDept.UpdatedAt),
		Children:    nil,
	}
	return resp, nil
}
