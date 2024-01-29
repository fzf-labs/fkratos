package data

import (
	"fkratos/app/rpc_common/internal/data/gorm/fkratos_common_repo"
	"fmt"

	conf "github.com/fzf-labs/fkratos-contrib/api/conf/v1"
	fcache "github.com/fzf-labs/fkratos-contrib/cache"
	"github.com/fzf-labs/fkratos-contrib/db"

	"github.com/fzf-labs/fpkg/orm/gen/cache"
	"github.com/fzf-labs/fpkg/orm/gen/cache/rueidisdbcache"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/redis/rueidis"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	db.NewGorm,
	fcache.NewRueidis,
	NewData,
	NewDBCache,
	NewSensitiveWordRepo,
	fkratos_common_repo.NewSensitiveWordRepo,
)

type Data struct {
	logger         *log.Helper
	db             *gorm.DB
	rueidis        rueidis.Client
	rueidisdbcache *rueidisdbcache.Cache
}

func NewData(c *conf.Bootstrap, logger log.Logger, db *gorm.DB, rueidisClient rueidis.Client) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", fmt.Sprintf("%s/data", c.Name)))
	d := &Data{
		logger:         l,
		db:             db,
		rueidis:        rueidisClient,
		rueidisdbcache: rueidisdbcache.NewRueidisDBCache(rueidisClient),
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
func NewDBCache(rueidisClient rueidis.Client) cache.IDBCache {
	return rueidisdbcache.NewRueidisDBCache(rueidisClient)
}
