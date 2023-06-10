package micro

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/micro/header"
	"github.com/metadiv-io/micro/system"
	"github.com/metadiv-io/micro/types"
)

func GetLogPrefix(ctx *gin.Context) string {
	traceID := header.GetTraceID(ctx)
	var apiUUID string
	api, ok := types.API_MAP[ctx.Request.Method+":"+ctx.FullPath()]
	if ok {
		apiUUID = api.Tag
	}
	return fmt.Sprintf("[system:%s] [api:%s] [trace:%s] [ip:%s]", system.SYSTEM_UUID, apiUUID, traceID, ctx.ClientIP())
}
