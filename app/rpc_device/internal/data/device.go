package data

import (
	"fkratos/app/rpc_device/internal/biz"
	"fkratos/app/rpc_device/internal/data/gorm/fkratos_device_repo"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.DeviceRepo = (*DeviceRepo)(nil)

func NewDeviceRepo(
	logger log.Logger,
	data *Data,
	deviceRepo *fkratos_device_repo.DeviceRepo,
) biz.DeviceRepo {
	l := log.NewHelper(log.With(logger, "module", "data/device"))
	return &DeviceRepo{
		log:        l,
		data:       data,
		DeviceRepo: deviceRepo,
	}
}

type DeviceRepo struct {
	log  *log.Helper
	data *Data
	*fkratos_device_repo.DeviceRepo
}
