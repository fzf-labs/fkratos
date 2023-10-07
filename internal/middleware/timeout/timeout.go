package timeout

import (
	"context"
	"errors"
	"fkratos/internal/errorx"

	"github.com/go-kratos/kratos/v2/middleware"
)

func Timeout() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			panicChan := make(chan any, 1)
			done := make(chan struct{})
			go func() {
				defer func() {
					if p := recover(); p != nil {
						panicChan <- p
					}
				}()
				reply, err = handler(ctx, req)
				close(done)
			}()
			select {
			case p := <-panicChan:
				panic(p)
			case <-ctx.Done():
				if errors.Is(ctx.Err(), context.Canceled) {
					return nil, errorx.RequestCanceledErr.WithMetadata(errorx.SetErrMetadata(ctx.Err())).WithCause(ctx.Err())
				}
				return nil, ctx.Err()
			case <-done:
				return reply, err
			}
		}
	}
}
