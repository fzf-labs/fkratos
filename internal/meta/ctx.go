package meta

import (
	"context"

	"github.com/fzf-labs/fpkg/conv"
)

type ContextWithValueKey string

var (
	ContextAdminId = "adminId"
)

// GetAdminId  从Context中获取adminId
func GetAdminId(ctx context.Context) string {
	return conv.String(ctx.Value(ContextWithValueKey(ContextAdminId)))
}

// SetAdminId 设置AdminId到Context
func SetAdminId(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, ContextWithValueKey(ContextAdminId), value)
}
