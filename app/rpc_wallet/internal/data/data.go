package data

import (
	"fkratos/internal/bootstrap"
	"fkratos/internal/bootstrap/conf"
	"fmt"

	"github.com/fzf-labs/fpkg/orm/gen/cache"
	"github.com/fzf-labs/fpkg/orm/gen/cache/rueidisdbcache"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/redis/rueidis"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	bootstrap.NewGorm,
	bootstrap.NewRueidis,
	NewDBCache,

	NewPayRepo,
	NewWalletRepo,
)

type Data struct {
	logger  *log.Helper
	gorm    *gorm.DB
	rueidis rueidis.Client
	dbCache cache.IDBCache
}

func NewData(c *conf.Bootstrap, logger log.Logger, db *gorm.DB, rueidisClient rueidis.Client, dbCache cache.IDBCache) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", fmt.Sprintf("%s/data", c.Name)))
	d := &Data{
		logger:  l,
		gorm:    db,
		rueidis: rueidisClient,
		dbCache: dbCache,
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
		d.rueidis.Close()
	}
	return d, cleanup, nil
}

func NewDBCache(rueidisClient rueidis.Client) cache.IDBCache {
	return rueidisdbcache.NewRueidisDBCache(rueidisClient)
}