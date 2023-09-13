package bootstrap

import (
	"fkratos/internal/bootstrap/conf"
	"fmt"

	"github.com/fzf-labs/fpkg/cache/gorediscache"
	"github.com/fzf-labs/fpkg/cache/rueidiscache"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
	"github.com/redis/rueidis"
)

// NewRedis 初始化Redis
func NewRedis(cfg *conf.Bootstrap, logger log.Logger) *redis.Client {
	l := log.NewHelper(log.With(logger, "module", "NewRedis"))
	r, err := gorediscache.NewGoRedis(gorediscache.GoRedisConfig{
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

// NewRueidis 初始化Rueidis
func NewRueidis(cfg *conf.Bootstrap, logger log.Logger) rueidis.Client {
	l := log.NewHelper(log.With(logger, "module", "NewRueidis"))
	r, err := rueidiscache.NewRueidis(&rueidis.ClientOption{
		Username:    cfg.Data.Redis.Username,
		Password:    cfg.Data.Redis.Password,
		InitAddress: []string{cfg.Data.Redis.Addr},
		SelectDB:    int(cfg.Data.Redis.Db),
	})
	if err != nil {
		l.Fatalf("failed opening connection to redis")
		panic(fmt.Sprintf("NewRueidis err: %s", err))
	}
	return r
}
