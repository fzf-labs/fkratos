package bootstrap

import (
	"fkratos/internal/bootstrap/conf"
	"fmt"

	"github.com/fzf-labs/fpkg/db"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

// NewGorm 初始化gorm
func NewGorm(cfg *conf.Bootstrap, logger log.Logger) *gorm.DB {
	l := log.NewHelper(log.With(logger, "module", "NewGorm"))
	client, err := db.NewGormPostgresClient(&db.GormPostgresClientConfig{
		DataSourceName:  cfg.Data.Gorm.DataSourceName,
		MaxIdleConn:     int(cfg.Data.Gorm.MaxIdleConn),
		MaxOpenConn:     int(cfg.Data.Gorm.MaxOpenConn),
		ConnMaxLifeTime: cfg.Data.Gorm.ConnMaxLifeTime.AsDuration(),
		ShowLog:         cfg.Data.Gorm.ShowLog,
		Tracing:         cfg.Data.Gorm.Tracing,
	})
	if err != nil {
		l.Fatalf("failed opening connection to postgres")
		panic(fmt.Sprintf("NewGorm postgres err: %s", err))
	}
	return client
}
