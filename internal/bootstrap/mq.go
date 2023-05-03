package bootstrap

import (
	"fkratos/internal/bootstrap/conf"
	"fkratos/pkg/asynq"
)

// NewAysnqClient 实例化Aysnq客户端(生产者)
func NewAysnqClient(cfg *conf.Bootstrap) *asynq.Client {
	return asynq.NewClient(asynq.NewRedisClientOpt(
		asynq.WithAddress(cfg.Data.Redis.GetAddr()),
		asynq.WithRedisAuth(cfg.Data.Redis.GetUsername(), cfg.Data.Redis.GetPassword()),
	))
}

// NewAysnqServer 实例化Aysnq服务(消费者)
func NewAysnqServer(cfg *conf.Bootstrap) *asynq.Server {
	return asynq.NewServer(asynq.NewRedisClientOpt(
		asynq.WithAddress(cfg.Data.Redis.GetAddr()),
		asynq.WithRedisAuth(cfg.Data.Redis.GetUsername(), cfg.Data.Redis.GetPassword()),
	))
}
