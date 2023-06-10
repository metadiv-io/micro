package auth

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/env"
	"github.com/metadiv-io/micro/call"
	"github.com/metadiv-io/micro/ginhelp"
	"github.com/metadiv-io/micro/header"
	"github.com/metadiv-io/micro/jwt"
	"github.com/metadiv-io/micro/system"
)

func AdminOnly(ctx *gin.Context) {
	if isMicro(ctx) {
		ctx.Next()
		return
	}
	claims := header.GetAuthClaims(ctx)
	if claims != nil && claims.Type == jwt.TYPE_ADMIN {
		if env.String("GIN_MODE", "") == "debug" { // skip IP check in debug mode
			ctx.Next()
			return
		}
		if claims.IP != ctx.ClientIP() && !isMicro(ctx) {
			ginhelp.AbortUnauthorized(ctx)
			return
		}
		ctx.Next()
		return
	}
	ginhelp.AbortUnauthorized(ctx)
}

func UserOnly(ctx *gin.Context) {
	claims := header.GetAuthClaims(ctx)
	if claims != nil {
		if claims.Type == jwt.TYPE_USER {
			workspace := header.GetWorkspace(ctx)
			if workspace != "" && claims.Workspace != workspace {
				ginhelp.AbortUnauthorized(ctx)
				return
			}
			if env.String("GIN_MODE", "") == "debug" { // skip IP check in debug mode
				ctx.Next()
				return
			}
			if claims.IP != ctx.ClientIP() && !isMicro(ctx) {
				ginhelp.AbortUnauthorized(ctx)
				return
			}
			ctx.Next()
			return
		}
	}
	if isMicro(ctx) {
		ctx.Next()
		return
	}
	ginhelp.AbortUnauthorized(ctx)
}

type IsMicroRequest struct {
	IP string `json:"ip"`
}

type IsMicroResponse struct {
	Allowed bool `json:"allowed"`
}

var isMicroCache = make(map[string]bool)
var microCacheExpiry = time.Now().Add(5 * time.Minute)

func isMicro(ctx *gin.Context) bool {
	if env.String("GIN_MODE", "") == "debug" {
		return true
	}
	if time.Now().After(microCacheExpiry) {
		isMicroCache = make(map[string]bool)
		microCacheExpiry = time.Now().Add(5 * time.Minute)
	}
	if isMicroCache[ctx.ClientIP()] {
		return true
	}
	resp, err := call.POST[IsMicroResponse](ctx, system.AUTH_SERVICE_URL+"/micro", &IsMicroRequest{
		IP: ctx.ClientIP(),
	}, nil)
	if err != nil || resp == nil || !resp.Success {
		return false
	}
	if !resp.Data.Allowed {
		return false
	}
	isMicroCache[ctx.ClientIP()] = true
	return true
}
