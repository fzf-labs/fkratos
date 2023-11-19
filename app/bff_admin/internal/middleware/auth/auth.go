package auth

import (
	"context"
	bffAdminV1 "fkratos/api/bff_admin/v1"
	sysV1 "fkratos/api/rpc_sys/v1"
	"fkratos/internal/errorx"
	"fkratos/internal/meta"
	"fmt"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// WhiteListMatcher 路由白名单
func WhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]bool)
	whiteList[bffAdminV1.Sys_SysAuthLoginCaptcha_FullMethodName] = true
	whiteList[bffAdminV1.Sys_SysAuthLogin_FullMethodName] = true
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

func SelectorMiddleware(authClient sysV1.SysAuthClient) middleware.Middleware {
	return selector.Server(
		Auth(authClient),
	).Match(WhiteListMatcher()).Build()
}

func Auth(authClient sysV1.SysAuthClient) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := http.RequestFromServerContext(ctx); ok {
				// Do something on entering
				// 获取header头部中的Authorization的值
				authorization := tr.Header.Get("Authorization")
				// 不存在则报错
				if authorization == "" {
					return nil, errorx.TokenNotRequest.Err()
				}
				// token截取
				var token string
				_, err := fmt.Sscanf(authorization, "Bearer %s", &token)
				if err != nil {
					return nil, errorx.TokenFormatErr.Err()
				}
				tokenClaims, err := authClient.SysAuthJwtTokenCheck(ctx, &sysV1.SysAuthJwtTokenCheckReq{Token: token})
				if err != nil {
					return nil, err
				}
				adminID := tokenClaims.AdminId
				// 将JwtUID参数写进context中
				ctx = meta.SetMetadata(ctx, meta.XMdAdminID, adminID)
			}
			return handler(ctx, req)
		}
	}
}
