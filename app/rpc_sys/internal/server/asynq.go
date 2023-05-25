package server

import (
	"context"
	"fkratos/app/rpc_sys/internal/data/mq"
	"fkratos/app/rpc_sys/internal/service"
	"fkratos/internal/bootstrap"
	"fkratos/internal/bootstrap/conf"
	asynq2 "fkratos/pkg/asynq"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/hibiken/asynq"
)

func NewAsynqServer(c *conf.Bootstrap, logger log.Logger, logService *service.OperationLogService) *asynq2.Server {
	l := log.NewHelper(log.With(logger, "module", "devices-stats/server/asynq"))
	srv := bootstrap.NewAysnqServer(c)
	err := srv.HandleFunc(mq.SysOperationLogStore, func(ctx context.Context, task *asynq.Task) error {
		return logService.SysOperationLogStoreConsumer(task.Payload())
	})
	if err != nil {
		l.Errorf("asynq HandleFunc err:%s", err)
		return nil
	}
	l.Info("asynq NewAsynqServer success")
	return srv
}
