package usage

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/micro"
	"github.com/metadiv-io/micro/auth"
	"github.com/metadiv-io/micro/call"
)

type UsageResponse struct {
	Allowed bool    `json:"allowed"`
	Credit  float64 `json:"credit"`
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
		resp, err := call.GET[UsageResponse](nil, USAGE_SERVICE_URL+"/usage", map[string]string{
			"workspace_uuid": workspaceUUID,
			"api_uuid":       apiUUID,
		}, nil)
		if err != nil || resp == nil || !resp.Success {
			return nil, err
		}
		if cachedUsage[workspaceUUID] == nil {
			cachedUsage[workspaceUUID] = make(map[string]*UsageResponse)
		}
		cachedUsage[workspaceUUID][apiUUID] = resp.Data
		api := micro.API_MAP[ctx.Request.Method+":"+ctx.FullPath()]
		api.Credit = resp.Data.Credit
		micro.API_MAP[ctx.Request.Method+":"+ctx.FullPath()] = api
		return resp.Data, nil
	}

	return cachedUsage[workspaceUUID][apiUUID], nil
}

func UsageRequired(ctx *gin.Context) {
	claims := auth.GetAuthClaims(ctx)
	if claims == nil {
		auth.AbortUnauthorized(ctx)
		return
	}
	workspace := micro.GetWorkspace(ctx)
	if workspace == "" {
		auth.AbortForbidden(ctx)
		return
	}
	apiUUID := micro.GetApiUUID(ctx)
	if apiUUID == "" {
		auth.AbortForbidden(ctx)
		return
	}
	usage, err := queryUsage(ctx, workspace, apiUUID)
	if err != nil || usage == nil {
		auth.AbortForbidden(ctx)
		return
	}
	addConsumption(workspace, usage.Credit)
	ctx.Next()
}
