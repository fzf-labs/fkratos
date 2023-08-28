package meta

import (
	"context"

	"github.com/go-kratos/kratos/v2/metadata"
)

var (
	XMdAdminID   = "x-md-global-adminId"
	XMdIP        = "x-md-global-ip"
	XMdUseragent = "x-md-global-useragent"
)

// GetMetadataFromServer 获取元数据-服务端Context
func GetMetadataFromServer(ctx context.Context, key string) string {
	if md, ok := metadata.FromServerContext(ctx); ok {
		return md.Get(key)
	}
	return ""
}

// GetMetadataFromClient 获取元数据-客户端Context
func GetMetadataFromClient(ctx context.Context, key string) string {
	if md, ok := metadata.FromClientContext(ctx); ok {
		return md.Get(key)
	}
	return ""
}

// SetMetadata 设置元数据
func SetMetadata(ctx context.Context, key, value string) context.Context {
	return metadata.AppendToClientContext(ctx, key, value)
}

// GetAdminIDFromClient 获取adminId-服务端Context
func GetAdminIDFromClient(ctx context.Context) string {
	return GetMetadataFromClient(ctx, XMdAdminID)
}

// GetAdminIDFromServer 获取adminId-客户端Context
func GetAdminIDFromServer(ctx context.Context) string {
	return GetMetadataFromServer(ctx, XMdAdminID)
}
