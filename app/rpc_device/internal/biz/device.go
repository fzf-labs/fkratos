package biz

import (
	"fkratos/app/rpc_device/internal/data/gorm/fkratos_device_repo"

	"github.com/go-kratos/kratos/v2/log"
)

type DeviceRepo interface {
	fkratos_device_repo.IDeviceRepo
}

func NewDeviceUseCase(
	logger log.Logger,
	deviceRepo DeviceRepo,
) *DeviceUseCase {
	l := log.NewHelper(log.With(logger, "module", "biz/device"))
	return &DeviceUseCase{
		log:        l,
		deviceRepo: deviceRepo,
	}
}

type DeviceUseCase struct {
	log        *log.Helper
	deviceRepo DeviceRepo
}
