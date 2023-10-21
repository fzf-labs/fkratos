package service

import (
	"context"

	pb "fkratos/api/rpc_device/v1"
)

func (d *DeviceService) DeviceOne(ctx context.Context, req *pb.DeviceOneReq) (*pb.DeviceOneReply, error) {
	return d.deviceUseCase.DeviceOne(ctx, req)
}
