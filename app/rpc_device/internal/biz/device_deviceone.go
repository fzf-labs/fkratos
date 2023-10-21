package biz

import (
	"context"
	"fkratos/internal/dto"
	"fkratos/internal/errorx"

	pb "fkratos/api/rpc_device/v1"
)

// DeviceOne 设备表-单条数据查询
func (d *DeviceUseCase) DeviceOne(ctx context.Context, req *pb.DeviceOneReq) (*pb.DeviceOneReply, error) {
	resp := &pb.DeviceOneReply{}
	result, err := d.deviceRepo.FindOneCacheByID(ctx, req.GetID())
	if err != nil {
		return nil, errorx.DataSQLErr.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	err = dto.Copy(&resp.Info, result)
	if err != nil {
		return nil, errorx.DataFormattingError.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return resp, nil
}
