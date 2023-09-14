package service

import (
	"context"
	rpcDeviceV1 "fkratos/api/rpc_device/v1"
	"fkratos/app/rpc_device/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

func NewDeviceService(logger log.Logger, deviceUseCase *biz.DeviceUseCase) *DeviceService {
	l := log.NewHelper(log.With(logger, "module", "Device/service/Device-service"))
	return &DeviceService{
		log:           l,
		deviceUseCase: deviceUseCase,
	}
}

type DeviceService struct {
	rpcDeviceV1.UnimplementedDeviceServer
	log *log.Helper

	deviceUseCase *biz.DeviceUseCase
}

func (u *DeviceService) CreateDevice(ctx context.Context, req *rpcDeviceV1.CreateDeviceReq) (*rpcDeviceV1.CreateDeviceReply, error) {
	//TODO implement me
	panic("implement me")
}

func (u *DeviceService) UpdateDevice(ctx context.Context, req *rpcDeviceV1.UpdateDeviceReq) (*rpcDeviceV1.UpdateDeviceReply, error) {
	//TODO implement me
	panic("implement me")
}

func (u *DeviceService) DeleteDevice(ctx context.Context, req *rpcDeviceV1.DeleteDeviceReq) (*rpcDeviceV1.DeleteDeviceReply, error) {
	//TODO implement me
	panic("implement me")
}

func (u *DeviceService) GetDevice(ctx context.Context, req *rpcDeviceV1.GetDeviceReq) (*rpcDeviceV1.GetDeviceReply, error) {
	//TODO implement me
	panic("implement me")
}

func (u *DeviceService) ListDevice(ctx context.Context, req *rpcDeviceV1.ListDeviceReq) (*rpcDeviceV1.ListDeviceReply, error) {
	//TODO implement me
	panic("implement me")
}
