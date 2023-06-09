package auth

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/env"
	"github.com/metadiv-io/micro"
	"github.com/metadiv-io/micro/call"
)

func AdminOnly(ctx *gin.Context) {
	if isMicro(ctx) {
		ctx.Next()
		return
	}
	claims := GetAuthClaims(ctx)
	if claims != nil && claims.Type == JWT_TYPE_ADMIN {
		if env.String("GIN_MODE", "") == "debug" { // skip IP check in debug mode
			ctx.Next()
			return
		}
		if !claims.HasIP(ctx.ClientIP()) && !isMicro(ctx) {
			AbortUnauthorized(ctx)
			return
		}
		ctx.Next()
		return
	}
	AbortUnauthorized(ctx)
}

func UserOnly(ctx *gin.Context) {
	claims := GetAuthClaims(ctx)
	if claims != nil {
		if claims.Type == JWT_TYPE_USER || claims.Type == JWT_TYPE_API || claims.Type == JWT_TYPE_WORKSPACE_USER {
			workspace := micro.GetWorkspace(ctx)
			if workspace != "" && !claims.HasWorkspace(workspace) {
				AbortUnauthorized(ctx)
				return
			}
			if env.String("GIN_MODE", "") == "debug" { // skip IP check in debug mode
				ctx.Next()
				return
			}
			if !claims.HasIP(ctx.ClientIP()) && !isMicro(ctx) {
				AbortUnauthorized(ctx)
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
	AbortUnauthorized(ctx)
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
	resp, err := call.POST[IsMicroResponse](ctx, AUTH_SERVICE_URL+"/micro", &IsMicroRequest{
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
