package data

import (
	"fkratos/internal/bootstrap"
	"fkratos/internal/bootstrap/conf"
	"fmt"

	"github.com/dtm-labs/rockscache"
	rc "github.com/fzf-labs/fpkg/cache/rockscache"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	bootstrap.NewGorm,
	bootstrap.NewRedis,
	rc.NewWeakRocksCacheClient,
	NewData,
	NewSensitiveWordRepo,
)

type Data struct {
	logger     *log.Helper
	db         *gorm.DB
	redis      *redis.Client
	rocksCache *rockscache.Client
}

func NewData(c *conf.Bootstrap, logger log.Logger, db *gorm.DB, redisClient *redis.Client, rocksCache *rockscache.Client) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", fmt.Sprintf("%s/data", c.ServiceName)))
	d := &Data{
		logger:     l,
		db:         db,
		redis:      redisClient,
		rocksCache: rocksCache,
	}
	cleanup := func() {
		log.Info("closing the data resources")
		sqlDB, err := d.db.DB()
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
