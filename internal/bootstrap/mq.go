package bootstrap

import (
	"fkratos/internal/bootstrap/conf"
	"fkratos/pkg/asynq"
)

// NewAysnq 实例化Aysnq
func NewAysnq(cfg *conf.Bootstrap) *asynq.Server {
	return asynq.NewServer(
		asynq.WithAddress(cfg.Data.Redis.GetAddr()),
		asynq.WithRedisAuth(cfg.Data.Redis.GetUsername(), cfg.Data.Redis.GetPassword()),
	)
}
