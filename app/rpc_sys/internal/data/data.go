package data

import (
	"fkratos/internal/bootstrap"
	"fkratos/internal/bootstrap/conf"
	"fkratos/internal/pkg/asynq"
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
	bootstrap.NewGorm,
	bootstrap.NewRueidis,
	bootstrap.NewAysnqClient,
	NewSysAdminRepo,
	NewSysDeptRepo,
	NewSysJobRepo,
	NewSysRoleRepo,
	NewSysLogRepo,
	NewSysAPIRepo,
	NewSysPermissionRepo,
)

type Data struct {
	logger         *log.Helper
	gorm           *gorm.DB
	rueidis        rueidis.Client
	rueidisdbcache *rueidisdbcache.Cache
	aysnqClient    *asynq.Client
}

func NewData(c *conf.Bootstrap, logger log.Logger, db *gorm.DB, rueidis rueidis.Client, aysnqClient *asynq.Client) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", fmt.Sprintf("%s/data", c.ServiceName)))
	d := &Data{
		logger:         l,
		gorm:           db,
		rueidis:        rueidis,
		rueidisdbcache: rueidisdbcache.NewRueidisDBCache(rueidis),
		aysnqClient:    aysnqClient,
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
