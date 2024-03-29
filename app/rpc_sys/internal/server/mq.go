package server

import (
	"context"
	"fkratos/app/rpc_sys/internal/data/mq"
	"fkratos/app/rpc_sys/internal/service"

	conf "github.com/fzf-labs/fkratos-contrib/api/conf/v1"

	"github.com/fzf-labs/fkratos-contrib/bootstrap"
	fasynq "github.com/fzf-labs/fkratos-contrib/mq/asynq"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/hibiken/asynq"
)

func NewAsynqServer(c *conf.Bootstrap, logger log.Logger, logService *service.SysLogService) *fasynq.Server {
	l := log.NewHelper(log.With(logger, "module", "devices-stats/server/asynq"))
	srv := bootstrap.NewAysnqServer(c)
	err := srv.HandleFunc(mq.SysLogStore, func(ctx context.Context, task *asynq.Task) error {
		return logService.SysLogStoreConsumer(task.Payload())
	})
	if err != nil {
		l.Errorf("asynq HandleFunc err:%s", err)
		return nil
	}
	l.Info("asynq NewAsynqServer success")
	return srv
}
