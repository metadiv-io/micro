package usage

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/logger"
	"github.com/metadiv-io/micro/call"
	"github.com/metadiv-io/micro/ginhelp"
	"github.com/metadiv-io/micro/header"
	"github.com/metadiv-io/micro/system"
	"github.com/metadiv-io/micro/types"
)

type UsageResponse struct {
	Allowed          bool    `json:"allowed"`
	Credit           float64 `json:"credit"`
	SubscriptionUUID string  `json:"subscription_uuid"`
}

type Usage struct {
	WorkspaceUUID string `json:"user_uuid"`
	ApiUUID       string `json:"api_uuid"`
}

var cachedUsage = make(map[string]map[string]*UsageResponse)
var clearCachedUsageAt time.Time

func queryUsage(ctx *gin.Context, workspaceUUID, apiUUID string) (*UsageResponse, error) {
	if clearCachedUsageAt.Before(time.Now()) {
		cachedUsage = make(map[string]map[string]*UsageResponse)
		clearCachedUsageAt = time.Now().Add(15 * time.Minute)
	}

	var askUsage bool
	_, ok1 := cachedUsage[workspaceUUID]
	if !ok1 {
		askUsage = true
	} else {
		_, ok2 := cachedUsage[workspaceUUID][apiUUID]
		if !ok2 {
			askUsage = true
		}
	}

	if askUsage {
		resp, err := call.GET[UsageResponse](nil, system.USAGE_SERVICE_URL+"/usage", map[string]string{
			"workspace_uuid": workspaceUUID,
			"api_uuid":       apiUUID,
		}, nil)
		if err != nil || resp == nil || !resp.Success {
			return nil, err
		}
		if !resp.Data.Allowed {
			return resp.Data, nil // return not allowed, no need to cache
		}
		if cachedUsage[workspaceUUID] == nil {
			cachedUsage[workspaceUUID] = make(map[string]*UsageResponse)
		}
		cachedUsage[workspaceUUID][apiUUID] = resp.Data
		api := types.API_MAP[ctx.Request.Method+":"+ctx.FullPath()]
		api.Credit = resp.Data.Credit
		types.API_MAP[ctx.Request.Method+":"+ctx.FullPath()] = api
		return resp.Data, nil
	}

	return cachedUsage[workspaceUUID][apiUUID], nil
}

func UsageRequired(ctx *gin.Context) {
	claims := header.GetAuthClaims(ctx)
	if claims == nil {
		ginhelp.AbortUnauthorized(ctx)
		return
	}
	workspace := header.GetWorkspace(ctx)
	if workspace == "" {
		ginhelp.AbortWorkspaceNotFound(ctx)
		return
	}
	apiUUID := header.GetApiUUID(ctx)
	if apiUUID == "" {
		ginhelp.AbortApiUUIDNotFound(ctx)
		return
	}
	usage, err := queryUsage(ctx, workspace, apiUUID)
	if err != nil {
		logger.Error("query usage:", err.Error())
		ginhelp.AbortForbidden(ctx)
		return
	}
	if !usage.Allowed {
		ginhelp.AbortNotEnoughCredit(ctx)
		return
	}
	addConsumption(usage.SubscriptionUUID, usage.Credit)
	ctx.Next()
}
