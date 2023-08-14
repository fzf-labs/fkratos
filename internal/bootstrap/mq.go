package bootstrap

import (
	"fkratos/internal/bootstrap/conf"
	asynq2 "fkratos/internal/pkg/asynq"
)

// NewAysnqClient 实例化Aysnq客户端(生产者)
func NewAysnqClient(cfg *conf.Bootstrap) *asynq2.Client {
	return asynq2.NewClient(asynq2.NewRedisClientOpt(
		asynq2.WithAddress(cfg.Data.Redis.GetAddr()),
		asynq2.WithRedisAuth(cfg.Data.Redis.GetUsername(), cfg.Data.Redis.GetPassword()),
	))
}

// NewAysnqServer 实例化Aysnq服务(消费者)
func NewAysnqServer(cfg *conf.Bootstrap) *asynq2.Server {
	return asynq2.NewServer(asynq2.NewRedisClientOpt(
		asynq2.WithAddress(cfg.Data.Redis.GetAddr()),
		asynq2.WithRedisAuth(cfg.Data.Redis.GetUsername(), cfg.Data.Redis.GetPassword()),
	))
}
