package biz

import (
	"context"
	pb "fkratos/api/rpc_sys/v1"

	"github.com/fzf-labs/fpkg/util/timeutil"
)

// SysJobInfo 岗位-单个岗位信息
func (s *SysJobUseCase) SysJobInfo(ctx context.Context, req *pb.SysJobInfoReq) (*pb.SysJobInfoReply, error) {
	resp := &pb.SysJobInfoReply{}
	sysJob, err := s.sysJobRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	resp.Info = &pb.SysJobInfo{
		Id:        sysJob.ID,
		Name:      sysJob.Name,
		Code:      sysJob.Code,
		Remark:    sysJob.Remark,
		Status:    int32(sysJob.Status),
		Sort:      int32(sysJob.Sort),
		CreatedAt: timeutil.ToDateTimeStringByTime(sysJob.CreatedAt),
		UpdatedAt: timeutil.ToDateTimeStringByTime(sysJob.UpdatedAt),
	}
	return resp, nil
}
