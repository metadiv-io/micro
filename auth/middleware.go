package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/micro"
	"github.com/metadiv-io/micro/jwt"
)

func SystemAdminOnly(ctx *gin.Context) {
	claims := GetClaims(ctx)
	if claims == nil {
		AbortUnauthorized(ctx)
		return
	}
	if claims.TokenType != jwt.TOKEN_TYPE_ADMIN_TOKEN && claims.TokenType != jwt.TOKEN_TYPE_SYSTEM_TOKEN {
		AbortUnauthorized(ctx)
		return
	}
}

func LoginRequired(ctx *gin.Context) {
	claims := GetClaims(ctx)
	if claims == nil {
		AbortUnauthorized(ctx)
		return
	}
	if claims.TokenType != jwt.TOKEN_TYPE_ACCESS_TOKEN && claims.TokenType != jwt.TOKEN_TYPE_API_TOKEN {
		AbortUnauthorized(ctx)
		return
	}
	ctx.Next()
}

func UsageRequired(method, path string) gin.HandlerFunc {
	micro.API_MAP[method+":"+path] = micro.Api{
		Tag: method + ":" + path,
	}
	return func(ctx *gin.Context) {
		claims := GetClaims(ctx)
		if claims == nil {
			AbortUnauthorized(ctx)
			return
		}
		if claims.TokenType != jwt.TOKEN_TYPE_ACCESS_TOKEN && claims.TokenType != jwt.TOKEN_TYPE_API_TOKEN {
			AbortUnauthorized(ctx)
			return
		}
		api, ok := micro.API_MAP[method+":"+path]
		if !ok {
			AbortForbidden(ctx)
			return
		}
		isUsageAllowed := isUsageAllowed(api.Tag, claims.Workspaces)
		if !isUsageAllowed {
			AbortForbidden(ctx)
			return
		}
		ctx.Next()
	}
}

func RefreshTokenOnly(ctx *gin.Context) {
	claims := GetClaims(ctx)
	if claims == nil {
		AbortUnauthorized(ctx)
		return
	}
	if claims.TokenType != jwt.TOKEN_TYPE_REFRESH_TOKEN {
		AbortUnauthorized(ctx)
		return
	}
}
