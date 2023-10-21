package service

import (
	"context"

	pb "fkratos/api/rpc_device/v1"
)

func (d *DeviceService) DeviceList(ctx context.Context, req *pb.DeviceListReq) (*pb.DeviceListReply, error) {
	return d.deviceUseCase.DeviceList(ctx, req)
}
