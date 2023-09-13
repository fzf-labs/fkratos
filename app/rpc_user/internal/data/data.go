package data

import (
	"fkratos/internal/bootstrap"
	"fkratos/internal/bootstrap/conf"
	"fmt"

	"github.com/fzf-labs/fpkg/orm/gen/cache/rueidisdbcache"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/redis/rueidis"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewUserRepo,
	bootstrap.NewGorm,
	bootstrap.NewRueidis,
)

type Data struct {
	logger         *log.Helper
	db             *gorm.DB
	rueidis        rueidis.Client
	rueidisdbcache *rueidisdbcache.Cache
}

func NewData(c *conf.Bootstrap, logger log.Logger, db *gorm.DB, rueidis rueidis.Client) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", fmt.Sprintf("%s/data", c.ServiceName)))
	d := &Data{
		logger:         l,
		db:             db,
		rueidis:        rueidis,
		rueidisdbcache: rueidisdbcache.NewRueidisDBCache(rueidis),
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
		d.rueidis.Close()
	}
	return d, cleanup, nil
}
