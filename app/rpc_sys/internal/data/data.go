package data

import (
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_repo"
	"fkratos/pkg/asynq"
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
	fkratos_sys_repo.NewSysAdminRepo,
	fkratos_sys_repo.NewSysDeptRepo,
	fkratos_sys_repo.NewSysJobRepo,
	fkratos_sys_repo.NewSysRoleRepo,
	fkratos_sys_repo.NewSysLogRepo,
	fkratos_sys_repo.NewSysAPIRepo,
	fkratos_sys_repo.NewSysPermissionRepo,
	NewSysAdminRepo,
	NewSysDeptRepo,
	NewSysJobRepo,
	NewSysRoleRepo,
	NewSysLogRepo,
	NewSysAPIRepo,
	NewSysPermissionRepo,
)

type Data struct {
	logger      *log.Helper
	gorm        *gorm.DB
	rueidis     rueidis.Client
	dbCache     cache.IDBCache
	aysnqClient *asynq.Client
}

func NewData(c *conf.Bootstrap, logger log.Logger, db *gorm.DB, dbCache cache.IDBCache, rueidisClient rueidis.Client, aysnqClient *asynq.Client) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", fmt.Sprintf("%s/data", c.Name)))
	d := &Data{
		logger:      l,
		gorm:        db,
		rueidis:     rueidisClient,
		dbCache:     dbCache,
		aysnqClient: aysnqClient,
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
