package header

import (
	"encoding/json"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/metadiv-io/micro/constant"
	"github.com/metadiv-io/micro/jwt"
	"github.com/metadiv-io/micro/system"
	"github.com/metadiv-io/micro/types"
)

func GetTraceID(c *gin.Context) string {
	traceID := c.GetHeader(constant.MICRO_HEADER_TRACE_ID)
	if traceID == "" {
		return uuid.NewString()
	}
	return traceID
}

func SetTraceID(c *gin.Context, traceID string) {
	c.Request.Header.Set(constant.MICRO_HEADER_TRACE_ID, traceID)
}

func GetTraces(c *gin.Context) []types.Trace {
	var traces []types.Trace
	traceHeader := c.GetHeader(constant.MICRO_HEADER_TRACES)
	if traceHeader != "" {
		_ = json.Unmarshal([]byte(traceHeader), &traces)
	}
	if len(traces) == 0 {
		traces = make([]types.Trace, 0)
	}
	return traces
}

func SetTraces(c *gin.Context, traces []types.Trace) {
	b, _ := json.Marshal(traces)
	c.Request.Header.Set(constant.MICRO_HEADER_TRACES, string(b))
}

func GetWorkspace(c *gin.Context) string {
	return c.GetHeader(constant.MICRO_HEADER_WORKSPACE)
}

func SetWorkspace(c *gin.Context, workspace string) {
	c.Request.Header.Set(constant.MICRO_HEADER_WORKSPACE, workspace)
}

func GetApiUUID(c *gin.Context) string {
	info, ok := types.API_MAP[c.Request.Method+":"+c.FullPath()]
	if ok {
		return info.UUID
	}
	return ""
}

func GetAuthClaims(ctx *gin.Context) *jwt.Claims {
	token := GetAuthToken(ctx)
	if token == "" {
		return nil
	}
	claims, err := jwt.ParseToken(token, system.JWT_PUBLIC_PEM)
	if err != nil {
		return nil
	}
	return claims
}

func GetAuthToken(ctx *gin.Context) string {
	t := ctx.GetHeader("Authorization")
	t = strings.ReplaceAll(t, "Bearer ", "")
	t = strings.ReplaceAll(t, " ", "")
	return t
}
