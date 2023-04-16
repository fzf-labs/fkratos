package bootstrap

import (
	"fkratos/internal/bootstrap/conf"
	"fmt"
	"github.com/fzf-labs/fpkg/cache"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
)

// NewRedis 初始化redis
func NewRedis(cfg *conf.Bootstrap, logger log.Logger) *redis.Client {
	l := log.NewHelper(log.With(logger, "module", "NewRedis"))
	r, err := cache.NewGoRedis(cache.GoRedisConfig{
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
