package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/micro"
	"github.com/metadiv-io/micro/jwt"
)

func SystemAdminOnly(ctx *gin.Context) {
	claims := GetClaims(ctx)
	if claims == nil {
		abortUnauthorized(ctx)
		return
	}
	if claims.TokenType != jwt.TOKEN_TYPE_ADMIN_TOKEN && claims.TokenType != jwt.TOKEN_TYPE_SYSTEM_TOKEN {
		abortUnauthorized(ctx)
		return
	}
}

func LoginRequired(ctx *gin.Context) {
	claims := GetClaims(ctx)
	if claims == nil {
		abortUnauthorized(ctx)
		return
	}
	if claims.TokenType != jwt.TOKEN_TYPE_ACCESS_TOKEN && claims.TokenType != jwt.TOKEN_TYPE_API_TOKEN {
		abortUnauthorized(ctx)
		return
	}
	isAllowed := isAllowed(claims.AuthMap)
	if !isAllowed {
		abortForbidden(ctx)
		return
	}
	api, ok := micro.API_MAP[ctx.Request.Method+":"+ctx.FullPath()]
	if !ok {
		abortForbidden(ctx)
		return
	}
	isUsageAllowed := isUsageAllowed(api.Tag, claims.AuthMap)
	if !isUsageAllowed {
		abortForbidden(ctx)
		return
	}
	ctx.Next()
}

func RefreshTokenOnly(ctx *gin.Context) {
	claims := GetClaims(ctx)
	if claims == nil {
		abortUnauthorized(ctx)
		return
	}
	if claims.TokenType != jwt.TOKEN_TYPE_REFRESH_TOKEN {
		abortUnauthorized(ctx)
		return
	}
}
