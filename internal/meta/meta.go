package meta

import (
	"context"

	"github.com/go-kratos/kratos/v2/metadata"
)

var (
	XMdAdminID   = "x-md-admin-id"
	XMdIP        = "x-md-ip"
	XMdUseragent = "x-md-useragent"
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

// GetAdminID 获取adminId
func GetAdminID(ctx context.Context) string {
	return GetMetadata(ctx, XMdAdminID)
}
