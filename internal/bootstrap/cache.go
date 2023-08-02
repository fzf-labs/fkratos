package bootstrap

import (
	"fkratos/internal/bootstrap/conf"
	"fmt"

	fRedis "github.com/fzf-labs/fpkg/cache/redis"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
)

// NewRedis 初始化redis
func NewRedis(cfg *conf.Bootstrap, logger log.Logger) *redis.Client {
	l := log.NewHelper(log.With(logger, "module", "NewRedis"))
	r, err := fRedis.NewGoRedis(fRedis.GoRedisConfig{
		Addr:         cfg.Data.Redis.Addr,
		Password:     cfg.Data.Redis.Password,
		DB:           int(cfg.Data.Redis.Db),
		DialTimeout:  cfg.Data.Redis.DialTimeout.AsDuration(),
		WriteTimeout: cfg.Data.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  cfg.Data.Redis.ReadTimeout.AsDuration(),
	})
	if err != nil {
		l.Fatalf("failed opening connection to redis")
		panic(fmt.Sprintf("NewRedis err: %s", err))
	}
	return r
}
