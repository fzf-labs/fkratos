package data

import (
	"fkratos/app/rpc_device/internal/biz"
	"fkratos/app/rpc_device/internal/data/gorm/fkratos_device_repo"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.DeviceRepo = (*DeviceRepo)(nil)

func NewDeviceRepo(data *Data, logger log.Logger) biz.DeviceRepo {
	l := log.NewHelper(log.With(logger, "module", "Device/data/Device"))
	return &DeviceRepo{
		data:       data,
		log:        l,
		DeviceRepo: fkratos_device_repo.NewDeviceRepo(data.gorm, data.rueidisdbcache),
	}
}

type DeviceRepo struct {
	data *Data
	log  *log.Helper
	*fkratos_device_repo.DeviceRepo
}

func (u *DeviceRepo) CreateDevice() {

}
