package service

import (
	pb "fkratos/api/rpc_device/v1"
	"fkratos/app/rpc_device/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

func NewDeviceService(
	logger log.Logger,
	deviceUseCase *biz.DeviceUseCase,
) *DeviceService {
	l := log.NewHelper(log.With(logger, "module", "service/device"))
	return &DeviceService{
		log:           l,
		deviceUseCase: deviceUseCase,
	}
}

type DeviceService struct {
	pb.UnimplementedDeviceServer
	log           *log.Helper
	deviceUseCase *biz.DeviceUseCase
}
