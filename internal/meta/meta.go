package meta

import (
	"context"

	"github.com/go-kratos/kratos/v2/metadata"
)

func GetMetadata(ctx context.Context, key string) string {
	if md, ok := metadata.FromServerContext(ctx); ok {
		return md.Get(key)
	}
	return ""
}

func SetMetadata(ctx context.Context, key, value string) context.Context {
	return metadata.AppendToClientContext(ctx, key, value)
}
