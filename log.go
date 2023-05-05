package micro

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetLogPrefix(ctx *gin.Context) string {
	traceID := GetTraceID(ctx)
	var apiUUID string
	api, ok := API_MAP[ctx.Request.Method+":"+ctx.FullPath()]
	if ok {
		apiUUID = api.Tag
	}
	return fmt.Sprintf("[%s/%s/%s@%s]", SYSTEM_UUID, apiUUID, traceID, ctx.ClientIP())
}
