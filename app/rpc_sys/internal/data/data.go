package data

import (
	"fkratos/internal/bootstrap"
	"fkratos/internal/bootstrap/conf"

	"github.com/dtm-labs/rockscache"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	bootstrap.NewGorm,
	bootstrap.NewRedis,
	bootstrap.NewRocksCache,

	NewSysAdminRepo,
	NewSysDeptRepo,
	NewSysJobRepo,
	NewSysRoleRepo,
	NewSysLogRepo,
	NewSysApiRepo,
)

// Data .
type Data struct {
	logger     *log.Helper
	gorm       *gorm.DB
	redis      *redis.Client
	rocksCache *rockscache.Client
}

// NewData .
func NewData(c *conf.Bootstrap, logger log.Logger, db *gorm.DB, redis *redis.Client, rocksCache *rockscache.Client) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "rpc_sys/data"))
	d := &Data{
		logger:     l,
		gorm:       db,
		redis:      redis,
		rocksCache: rocksCache,
	}
	cleanup := func() {
		log.Info("closing the data resources")
		sqlDB, err := d.gorm.DB()
		if err != nil {
			l.Error(err)
		}
		err = sqlDB.Close()
		if err != nil {
			l.Error(err)
		}
		err = d.redis.Close()
		if err != nil {
			l.Error(err)
		}
	}
	return d, cleanup, nil
}
